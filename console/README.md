# Console

**目录**

- [使用docker-compose快速开始](#使用docker-compose快速开始)

## 使用`docker-compose`快速开始

> [Compose命令说明](https://yeasy.gitbooks.io/docker_practice/content/compose/commands.html)

**服务列表:**
- 注册中心Etcd
    - `etcd`，使用docker镜像`bitnami/etcd`
- API网关
    - `gateway`，使用docker镜像`hbchen/starter-kit-gateway`
- Console Web
    - `web`
- Console API
    - `api`
- Account SRV
    - `account`

**启动Console全部服务**

```shell script
# 首次运行创建docker network
docker network create starter-kit-console
```

Go编译，编译console项目全部服务，包括`web`、`api`和`account`服务
```shell script
make build
```

启动服务
> `p`参数项目名称，可以自己定义
```shell script
make start p=starter-kit-console
# 或
docker-compose -p starter-kit-console up
```

**重新编译并启动服务**

> 如重启`api`服务

**方法一:**
使用Makefile的`make restart`命令，`p`项目名称，`s`服务名称
```shell script
 make restart p=starter-kit-console s=api
```

**方法二:**
使用docker-compose分步操作，`stop`停止服务、Go编译、`build`重新构建服务镜像、`up`启动服务
```shell script
docker-compose -p starter-kit-console stop api
cd api && make build_linux && cd ..
docker-compose -p starter-kit-console build api
docker-compose -p starter-kit-console up -d --no-deps --force-recreate api
```

## Proto管理
项目`.proto`文件统一在`pb`目录下，协议生成的`out`输出到指定子项目位置`gen=path`
```shell script
make proto gen=api/genproto
make proto gen=account/genproto
```

