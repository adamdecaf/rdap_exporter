VERSION := $(shell grep -Eo '(\d\.\d\.\d)(-dev)?' main.go | head -n1)

.PHONY: build deps docker

build:
	go fmt ./...
	go vet ./...
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/rdap-exporter-darwin .
	CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -o bin/rdap-exporter-linux .

deps:
	dep ensure -update

docker: build
	docker build -t adamdecaf/rdap-exporter:$(VERSION) -f Dockerfile .
