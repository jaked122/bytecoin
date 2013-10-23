all: bin/ytcd

bin/ytcd: src/ytcd/*.go src/libytcd/*.go
	rm -f bin/ytcd
	GOPATH=$(CURDIR) go install ytcd

test: all
	GOPATH=$(CURDIR) go test ytcd libytcd


.PHONY: all test