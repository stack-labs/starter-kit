
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. pb/account/*.proto
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. pb/hipstershop/*.proto
