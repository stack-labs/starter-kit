
GOPATH:=$(shell go env GOPATH)

PB_PATH=./pb
GEN_PATH=github.com/micro-in-cn/starter-kit/console/${gen}

# 协议生成，输出到指定项目位置gen=path
# make proto gen=api/genproto
# make proto gen=account/genproto
.PHONY: proto
proto:
	rm -rf ${gen}/*
	protoc --proto_path=${GOPATH}/src:${PB_PATH} \
	--micro_out=Msrv/account.proto=${GEN_PATH}/srv,Mgithub.com/micro/go-micro/api/proto/api.proto=github.com/micro/go-micro/v3/api/proto:${gen} \
	--go_out=Msrv/account.proto=${GEN_PATH}/srv,Mgithub.com/micro/go-micro/api/proto/api.proto=github.com/micro/go-micro/v3/api/proto:${gen} \
	--validate_out="lang=go:${gen}" \
	api/account.proto
	protoc --proto_path=${GOPATH}/src:${PB_PATH} \
	--micro_out=${gen} \
	--go_out=:${gen} \
	--validate_out="lang=go:${gen}" \
	srv/account.proto

# Go编译
.PHONY: build
build:
	cd account && make build_linux
	cd api && make build_linux
	cd web && make vue build_linux


COMPOSE_FILE=docker-compose.yml

# 启动服务，首次需要build
.PHONY: start
start:
	docker-compose -f ${COMPOSE_FILE} -p $(p) build --no-cache
	docker-compose --verbose -f ${COMPOSE_FILE} -p $(p) up

.PHONY: stop
stop:
	docker-compose --verbose -f ${COMPOSE_FILE} -p $(p) down

# 重新编译并启动服务
.PHONY: restart
restart:
	docker-compose -f ${COMPOSE_FILE} -p $(p) stop $(s)
	cd $(s) && make build_linux
	docker-compose -f ${COMPOSE_FILE} -p $(p) build --no-cache $(s)
	docker-compose -f ${COMPOSE_FILE} -p $(p) up -d --no-deps --force-recreate $(s)

.PHONY: restart_service
restart_service:
	docker-compose -f ${COMPOSE_FILE} -p $(p) stop $(s)
	docker-compose -f ${COMPOSE_FILE} -p $(p) up -d --no-deps --force-recreate $(s)
