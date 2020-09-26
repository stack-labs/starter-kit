
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	cd .. && make proto gen=account/genproto

.PHONY: build
build:
	go build -o ./bin/account-srv cmd/main.go

.PHONY: build_linux
build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -o ./bin/linux_amd64/account-srv cmd/main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build_linux
	docker build . -t $(tag)

.PHONY: run
run:
	./bin/account-srv --profile=$(profile) --conf_path=./conf/
