
GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	go build -o ./bin/micro main.go plugin.go

.PHONY: build_linux
build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -o ./bin/linux_amd64/micro main.go plugin.go

.PHONY: docker
docker: build_linux
	docker build . -t $(tag)

.PHONY: run_api
run_api:
	./bin/micro --registry=$(registry) --transport=$(transport) api

.PHONY: run_web
run_web:
	./bin/micro --registry=$(registry) --transport=$(transport) web
