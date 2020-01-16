
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	go build -o ./bin/console-web main.go plugin.go
	# go build -o ./bin/console-web -tags "statik" main.go plugin.go

.PHONY: build_linux
build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -o ./bin/linux_amd64/console-web main.go plugin.go
	# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -o ./bin/linux_amd64/console-web -tags "statik" main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: vue build_linux
	docker build . -t $(tag)

.PHONY: vue
vue:
	cd vue && npm install  && npm run build:prod

.PHONY: statik
statik:
	go get github.com/rakyll/statik
	go generate ./...

.PHONY: run
run:
	./bin/console-web --registry=$(registry) --transport=$(transport)
