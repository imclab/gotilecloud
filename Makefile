SHELL = bash

.PHONY: gofmt
gofmt:
	for file in $$(find src -name \*.go) ; do diff -u $$file <(gofmt $$file) ; done
