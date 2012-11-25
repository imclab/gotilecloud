SHELL = bash

all: gofmt test

.PHONY: gofmt
gofmt:
	for file in $$(find src -name \*.go) ; do diff -u $$file <(gofmt $$file) ; done

.PHONY: test
test:
	go test tilecloud/geom tilecloud/tile
