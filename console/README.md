# Console

- Go环境
    - `go 1.13`
    - `export GOSUMDB=off`
    - `export GOPROXY=https://mirrors.aliyun.com/goproxy/,direct`

## 目录

- [架构](#架构)
- [快速开始](#快速开始)
    - [Docker Compose启动](#docker-compose启动)`推荐`
    - [编译启动](#编译启动)
- [服务测试](#服务测试)
- [Proto管理](#proto管理)

## 架构
<img src="/doc/img/console-design.png" width="75%">

## 快速开始
### Docker Compose启动

使用`docker-compose`快速启动服务，适合搭建本地开发环境，省去每个服务去做启动管理的烦恼，Compose启动后可以对单个服务进行重新编译启动。

> [Compose命令参考](https://yeasy.gitbooks.io/docker_practice/content/compose/commands.html)

Compose包含以下服务:
- Etcd注册中心
    - `etcd`，使用docker镜像`bitnami/etcd`
- API网关
    - `gateway`，使用docker镜像`hbchen/starter-kit-gateway`
- Console Web
    - `web`
- Console API
    - `api`
- Account SRV
    - `account`

**Compose启动服务**

1.首次运行创建Docker网络
```shell script
docker network create starter-kit-console
```

2.Go编译，编译Console项目全部服务，包括`web`、`api`和`account`服务
> 使用Dockerfile编译太慢，所以编译还是选择使用本机环境
```shell script
make build
```

3.启动服务
> `p`参数为项目名称，可以自己定义
```shell script
make start p=starter-kit-console
# 或
docker-compose -p starter-kit-console up
```

4.[服务测试](#服务测试)

**重新编译并启动服务**

> 例如要重启`api`服务
> 注意：`web`服务前端没有自动编译，需要单独运行`make vue`更新前端

**方法一:**
使用`Makefile`的`make restart`命令，`p`项目名称，`s`服务名称
```shell script
 make restart p=starter-kit-console s=api
```

**方法二:**
使用`docker-compose`分步操作，`stop`停止服务、Go编译、`build`重新构建服务镜像、`up`启动服务
```shell script
docker-compose -p starter-kit-console stop api
cd api && make build_linux && cd ..
docker-compose -p starter-kit-console build api
docker-compose -p starter-kit-console up -d --no-deps --force-recreate api
```

**Compose常用命令**
```shell script
# 移除compose file中没有定义的容器
docker-compose -p starter-kit-console down --remove-orphans
```

### 编译启动

> 如果使用Etcd注册中心需要自己维护

**1.运行网关**

[网关](./../gateway) 

```bash
$ cd gateway

# 编译
$ make build

# API网关(二选一)
$ make run_api                                  # 默认mdns + http
$ make run_api registry=etcd transport=tcp      # 使用etcd + tcp
```

**2.运行服务**
- Web服务
	- `console/web`
- API服务
	- `console/api`
- Account服务
	- `console/account`
	
> 注：`registry`、`transport`选择与网关一致
```bash
$ cd {指定服务目录}

# 运行服务(二选一)
$ make build run                                # 默认mdns + http
$ make build run registry=etcd transport=tcp    # 使用etcd + tcp
```

**Makefile说明**
```bash
$ make build                                    # 编译
$ make run                                      # 运行
$ make run registry=etcd transport=tcp          # 运行，指定registry、transport

$ make build run                                # 编译&运行
$ make build run registry=etcd transport=tcp    # 编译&运行，指定registry、transport

$ make vue statik                               # 前端编译，并打包statik.go文件

$ make docker tag=xxx/xxx:v0.0.1
```

### 服务测试
> 注：`console API`由于有`认证`不能直接访问
- gateway
	- http://localhost:8080/
	- http://localhost:8080/metrics
- console
	- Web
	    - http://localhost:8080/console/
		- http://localhost:8080/console/v1/echo/
		- http://localhost:8080/console/v1/gin/
		- http://localhost:8080/console/v1/iris/
		- http://localhost:8080/console/v1/beego/
	- API
        - http://localhost:8080/account/login
        - http://localhost:8080/account/info
        - http://localhost:8080/account/logout
        
## Proto管理
项目`.proto`文件统一在`pb`目录下，协议生成的`out`输出到指定子项目位置`gen=path`
```shell script
# console目录运行make proto输出到子项目的genproto
make proto gen=api/genproto
make proto gen=account/genproto

# 或在子项目目录运行make proto，实际仍是调用console下的make proto
cd api
make proto
```

### Swagger生成
```shell script
cd pb
protoc -I$GOPATH/src/ -I./ \
--swagger_out=logtostderr=true,grpc_api_configuration=api/api.yaml,allow_merge=true,merge_file_name=api/api:. \
api/*.proto
```
