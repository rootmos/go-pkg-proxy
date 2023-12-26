package common

import (
	"os"
)

const (
	WhoAmI = "go-pkg-proxy"
	EnvPrefix = "GO_PKG_PROXY_"
)

func Getenv(key string) string {
	return os.Getenv(EnvPrefix + key)
}

func Getenv2(key string, def string) string {
	v, isset := os.LookupEnv(EnvPrefix + key)
	if isset {
		return v
	} else {
		return def
	}
}

func GetenvBool(key string) bool {
	return Getenv(key) != ""
}
