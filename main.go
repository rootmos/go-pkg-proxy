package main

import (
	"log"
	"net/http"
	"flag"
	"fmt"
	"strings"
	"encoding/json"
	"context"

	"rootmos.io/go-pkg-proxy/internal/common"
	"rootmos.io/go-pkg-proxy/internal/osext"
	"rootmos.io/go-pkg-proxy/internal/logging"
)

type Module struct {
	Name string `json:"name"`
	Root string `json:"root"`
	VCS string `json:"vcs,omitempty"`
	Repo string`json:"repo"`
}

type Modules map[string]Module

func FetchModules(ctx context.Context, url string) (Modules, error) {
	logger := logging.Get(ctx)
	f, err := osext.Open(ctx, url)
	defer f.Close()

	var raw []Module
	dec := json.NewDecoder(f)
	err = dec.Decode(&raw)
	if err != nil {
		return nil, err
	}

	modules := make(Modules)
	for _, m := range raw {
		logger.Debug("module", "name", m.Name, "definition", m)
		modules[m.Name] = m
	}

	return modules, nil
}

func main() {
	addr := flag.String("addr", common.Getenv2("ADDR", ":8000"), "bind to addr:port")
	modulesURL := flag.String("modules", common.Getenv2("MODULES", "file://go.json"), "fetch modules from URL")
	flag.Parse()

	logger, err := logging.SetupDefaultLogger()
	if err != nil {
		log.Fatal(err)
	}
	logger.Debug("hello")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger := logger.With("url", r.URL, "remoteAddr", r.RemoteAddr)
		logger.Info("request")

		ctx := logging.Set(r.Context(), logger)

		modules, err := FetchModules(ctx, *modulesURL)
		if err != nil {
			http.Error(w, "unable to fetch modules", http.StatusInternalServerError)
			return
		}

		modpath, _ := strings.CutPrefix(r.URL.Path, "/")
		mod, found := modules[modpath]

		write := func(str string) {
			if err != nil {
				return
			}
			_, err = w.Write([]byte(str))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		write("<html><head>")
		if found {
			vcs := mod.VCS
			if vcs == "" {
				vcs = "git"
			}
			write(fmt.Sprintf("<meta name=\"go-import\" content=\"%s %s %s\">", mod.Root, vcs, mod.Repo))
		}
		write("</head>")
		write("<body></body>")
		write("</html>")
	})

	logger.Info("listening", "addr", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
