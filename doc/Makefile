THIS := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
ROOT ?= ..
TOOLS ?= $(ROOT)/tools

export EXE ?= $(ROOT)/target/go-pkg-proxy
export APP ?= go-pkg-proxy

INCLUDE ?= $(TOOLS)/include
USAGE ?= $(TOOLS)/usage

.PHONY: all
all: $(ROOT)/README.md nginx.conf

$(ROOT)/README.md: README.in.md usage
	$(INCLUDE) "$<" "$@"

nginx.conf: nginx.in.conf nginx.snippet.conf
	$(INCLUDE) "$<" "$@"

usage: $(EXE)
	$(USAGE) $@

.PHONY: clean
clean:
	rm -f usage
