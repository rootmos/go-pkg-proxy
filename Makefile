EXE ?= go-pkg-proxy

TARGET ?= $(abspath ./target)

GO ?= go
export GOPATH ?= $(abspath ./go)

build: FORCE
	$(GO) build -v -o $(abspath $(TARGET))/ .

run: build
	$(TARGET)/$(EXE)

test:
	$(GO) test $(if $(VERBOSE),-v,) ./...

update:
	$(GO) get -u
	$(GO) mod tidy

doc: build
	$(MAKE) -C doc

clean:
	rm -rf $(TARGET)

deepclean: clean
	-chmod +w -R $(GOPATH)
	rm -rf $(GOPATH)

.PHONY: build run test
.PHONY: update doc
.PHONY: clean deepclean
FORCE:
