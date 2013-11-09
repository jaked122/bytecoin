all: bin/ytcd bin/ytcClient

bin/ytcd: src/ytcd/*.go src/libytcd/*.go
	rm -f bin/ytcd
	GOPATH=$(CURDIR) go install ytcd

bin/ytcClient: src/ytcClient/*.go src/libytcd/*.go
	rm -f bin/ytcClient
	GOPATH=$(CURDIR) go install ytcClient

test: all
	GOPATH=$(CURDIR) go test libGFC

fmt:
	go fmt src/ytcd/*.go
	go fmt src/libytcd/*.go
	go fmt src/ytcClient/*.go
	go fmt src/libGFC/*.go

clean:
	rm -f bin/ytcd
	rm -f bin/ytcClient


.PHONY: all test
