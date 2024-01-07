go-pkg-proxy
============

[![builds.sr.ht status](https://builds.sr.ht/~rootmos/go-pkg-proxy.svg)](https://builds.sr.ht/~rootmos/go-pkg-proxy?)

A tiny HTTP server that helps you configure repository URL:s for go modules,
using the [`?go-get=1`](https://go.dev/ref/mod#vcs-find) query parameter.
Useful when you have your own domain, but your modules live in different
version control systems.

## Example

Modules and the corresponding reposity URLs are configured in a JSON file:
```json
[
    {
        "name": "go-pkg-proxy",
        "root": "rootmos.io/go-pkg-proxy",
        "repo": "https://git.sr.ht/~rootmos/go-pkg-proxy"
    },
    {
        "names": [
            "go-utils",
            "go-utils/logging"
        ],
        "root": "rootmos.io/go-utils",
        "repo": "https://git.sr.ht/~rootmos/go-utils"
    }
]
```

and then configure your internet-facing HTTP server to route the `?go-get=1` requests.
For example using [nginx](https://nginx.org/en/docs/http/ngx_http_core_module.html#var_args):
```
location / {
  if ($args ~ "go-get=1") {
    proxy_pass http://127.0.0.1:8000;
    break;
  }

  root webroot;
}
```
(A complete configuration can be found [here](doc/nginx.conf) which is [tested here](test-nginx.py).)

## Usage
```
Usage of go-pkg-proxy:
  -addr string
    	bind to addr:port (default ":8000")
  -dry-run
    	try loading modules and exit afterwards
  -json-log-file string
    	log JSON to file (default "/dev/null")
  -json-log-level string
    	set JSON log level (default "INFO")
  -log-file string
    	log to file (default "/dev/stderr")
  -log-level string
    	set log level (default "INFO")
  -modules string
    	fetch modules from URL (default "file://go.json")
```
