
GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	go build -o micro main.go plugin.go

.PHONY: docker
docker:
	docker build . -t micro:latest

.PHONY: run
run:
	go run main.go plugin.go -h
