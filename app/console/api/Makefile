
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/*/*.proto

.PHONY: build
build: proto
	go build -o ./bin/console-api main.go plugin.go

.PHONY: build_linux
build_linux: proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -o ./bin/linux_amd64/console-api main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build_linux
	docker build . -t $(tag)

.PHONY: run
run:
	./bin/console-api --registry=$(registry)
