THIS := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
ROOT ?= ..
TOOLS ?= $(ROOT)/tools

export EXE ?= $(ROOT)/target/go-pkg-proxy
export APP ?= go-pkg-proxy

INCLUDE ?= $(TOOLS)/include
USAGE ?= $(TOOLS)/usage

.PHONY: all
all: $(ROOT)/README.md

$(ROOT)/README.md: README.in.md usage
	$(INCLUDE) "$<" "$@"

usage: $(EXE)
	$(USAGE) $@

.PHONY: clean
clean:
	rm -f usage
