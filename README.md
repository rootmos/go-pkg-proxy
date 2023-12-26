go-pkg-proxy
============

## Example

```json
[
    {
        "name": "go-pkg-proxy",
        "root": "rootmos.io/go-pkg-proxy",
        "repo": "https://git.sr.ht/~rootmos/go-pkg-proxy"
    }
]
```

```json
location / {
  if ($args ~ "go-get=1") {
    proxy_pass http://127.0.0.1:8000;
    break;
  }

  root webroot;

```

## Usage
```
Usage of go-pkg-proxy:
  -addr string
    	bind to addr:port (default ":8000")
  -dry-run
    	try loading modules and exit afterwards
  -log-json
    	log JSON objects (instead of plain-text)
  -log-level string
    	set logging level
  -modules string
    	fetch modules from URL (default "file://go.json")
```
