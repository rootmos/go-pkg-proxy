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
@include "go.json"
```

and then configure your internet-facing HTTP server to route the `?go-get=1` requests.
For example using [nginx](https://nginx.org/en/docs/http/ngx_http_core_module.html#var_args):
```
@include "nginx.snippet.conf"
```
(A complete configuration can be found [here](doc/nginx.conf) which is [tested here](test-nginx.py).)

## Usage
```
@include "usage"
```
