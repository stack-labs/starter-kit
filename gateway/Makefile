
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --proto_path=$GOPATH/src --stack_out=$GOPATH/src --go_out=:$GOPATH/src github.com/stack-labs/starter-kit/pkg/gateway/api/proto/api.proto

.PHONY: build
build:
	go build -o ./bin/micro main.go plugin.go

.PHONY: build_linux
build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -o ./bin/linux_amd64/micro main.go plugin.go

.PHONY: docker
docker: build_linux
	docker build . -t $(tag)

.PHONY: run
run:
	./bin/micro --config=conf/stack_config.yml
