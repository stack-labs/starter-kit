
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build: vue statik
	go build -o console-web main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t console-web:latest

.PHONY: vue
vue:
	cd vue && npm run build:prod

.PHONY: statik
statik:
	go get github.com/rakyll/statik
	go generate ./...

.PHONY: run
run: vue statik
	go run main.go
