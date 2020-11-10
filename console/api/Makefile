
GOPATH:=$(shell go env GOPATH)

PB_PATH=${GOPATH}/src/github.com/micro-in-cn/starter-kit/console/pb
GENPROTO=github.com/micro-in-cn/starter-kit/console/api/genproto

.PHONY: proto
proto:
	cd .. && make proto gen=api/genproto

.PHONY: build
build:
	go build -o ./bin/console-api main.go plugin.go

.PHONY: build_linux
build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -o ./bin/linux_amd64/console-api main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build_linux
	docker build . -t $(tag)

.PHONY: run
run:
	./bin/console-api --registry=$(registry) --transport=$(transport)
