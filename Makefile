all: bin/ytcd bin/ytcClient

bin/ytcd: libraries
	rm -f bin/ytcd
	GOPATH=$(CURDIR) go install ytcd

bin/ytcClient: libraries
	rm -f bin/ytcClient
	GOPATH=$(CURDIR) go install ytcClient

libraries: src/libGFC/*.go src/libytcd/*.go src/libytcd/*.go
	GOPATH=$(CURDIR) go install libGFC libytc libytcd

test: libraries
	GOPATH=$(CURDIR) go test libGFC libytc libytcd

fmt:
	go fmt src/ytcd/*.go
	go fmt src/ytcClient/*.go
	go fmt src/libytcd/*.go
	go fmt src/libGFC/*.go
	go fmt src/libytc/*.go

clean:
	rm -f bin/ytcd
	rm -f bin/ytcClient


.PHONY: all test fmt clean libraries
