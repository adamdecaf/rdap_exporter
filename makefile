VERSION := $(shell grep -Eo '(v[0-9]+[\.][0-9]+[\.][0-9]+([-a-zA-Z0-9]*)?)' main.go)

.PHONY: build deps docker

.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	go test ./...
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	COVER_THRESHOLD=10.0 ./lint-project.sh
endif

build:
	go fmt ./...
	go vet ./...
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/rdap-exporter-darwin .
	CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -o bin/rdap-exporter-linux .

docker: build
	docker build --pull -t adamdecaf/rdap_exporter:$(VERSION) -f Dockerfile .

release-push:
	docker push adamdecaf/rdap_exporter:$(VERSION)

test:
	go test -v ./...
