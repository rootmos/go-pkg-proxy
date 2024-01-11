package main

import (
	"net/http"
	"flag"
	"fmt"
	"strings"
	"encoding/json"
	"context"

	"rootmos.io/go-pkg-proxy/internal/common"
	"rootmos.io/go-utils/logging"
	"rootmos.io/go-utils/osext"
)

type Module struct {
	Name string `json:"name"`
	Names []string `json:"names"`
	Root string `json:"root"`
	VCS string `json:"vcs,omitempty"`
	Repo string`json:"repo"`
}

type Modules map[string]*Module

func FetchModules(ctx context.Context, url string) (Modules, error) {
	logger := logging.Get(ctx)
	f, err := osext.Open(ctx, url)
	defer f.Close()

	var raw []Module
	err = json.NewDecoder(f).Decode(&raw)
	if err != nil {
		return nil, err
	}

	modules := make(Modules)
	for _, m := range raw {
		mod := m
		if m.Name != "" {
			logger.Debug("module", "name", m.Name, "definition", m)
			modules[m.Name] = &mod
		}

		for _, n := range m.Names {
			logger.Debug("module", "name", n, "definition", m)
			modules[n] = &mod
		}
	}

	return modules, nil
}

func main() {
	addr := flag.String("addr", common.Getenv2("ADDR", ":8000"), "bind to addr:port")
	modulesURL := flag.String("modules", common.Getenv2("MODULES", "file://go.json"), "fetch modules from URL")
	dryRun := flag.Bool("dry-run", common.GetenvBool("DRY_RUN"), "try loading modules and exit afterwards")

	logConfig := logging.PrepareConfig(common.EnvPrefix)

	flag.Parse()

	logger, closer, err := logConfig.SetupDefaultLogger()
	if err != nil {
		logger.Exit(1, "unable to configure logger: %v", err)
	}
	defer closer()
	logger.Debug("hello")

	if *dryRun {
		ctx := logging.Set(context.Background(), logger)

		_, err := FetchModules(ctx, *modulesURL)
		if err != nil {
			logger.Exit(1, "unable to fetch modules: %v", err)
		}
		return
	}

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
	if err := http.ListenAndServe(*addr, nil); err != nil {
		logger.Exit(1, "serving failed: %v", err)
	}
}
