# Micro 快速开发工具包*项目进行中*

本仓库旨在提供面向Go-Micro生产环境的快速开发包。项目结合维护者们十余年的工作经验，不同领域的实战沉淀，一切为了缩短大家的选型、开发周期。

## 目录

- [目标](#目标)
- [架构设计](#架构设计)
- [快速开始](#快速开始)

## 目标

- 自定义[micro网关](gateway)
	- [x] `JWT`认证
	- [x] `Casbin`鉴权
	- [ ] Tracing
	- [ ] RequestID
	- [x] Metrics
	- [ ] Access Log
	- ...
- 网关选择
	- micro api
		- [x] endpoint
		- [ ] proxy
		- [ ] rpc
	- micro web
		- [ ] gin
		- [ ] echo
		- [x] [iris](/app/console/web/iris)
- 配置中心
- 前后端分离`console`
	- [x] [PanJiaChen/vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)
	- [ ] [tookit/vue-material-admin](https://github.com/tookit/vue-material-admin) 
	- [ ] [view-design/iview-admin](https://github.com/view-design/iview-admin)
- 参数验证
	- [ ] [protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate)
- 领域驱动
	- [x] 整洁架构
- ORM
	- [ ] gorm
	- [ ] xorm
- 发布
	- [ ] 灰度
	- [ ] 蓝绿
- 部署
	- [ ] K8S
		- [ ] helm
	- [ ] Docker
- 安全
- CICD
	- [ ] Jenkins
- 基础服务
	- [ ] 日志收集
		- `stdout`标准输出
		- `log.file`日志文件
		- [log-pilot](https://github.com/AliyunContainerService/log-pilot) 
	- [ ] 监控告警
		- Prometheus
		- Grafana
	- [ ] Tracing
		- Jaeger
- ...

## 架构设计

### 目录结构

```bash
├── app                 应用，API聚合、Web应用
│   ├── console         控制台
│   │   ├── api         go.micro.api.*，API
│   │   └── web         go.micro.web.*，Web，集成gin、echo、iris等web框架
│   ├── mobile          移动端
│   └── openapi         开放API
├── deploy              部署
│   ├── docker
│   └── k8s
├── doc                 文档资源
├── gateway             网关，自定义micro
├── pkg                 公共资源包
└── srv                 基础服务
    ├── account         账户服务，领域模型整洁架构示例
    │   ├── domain              领域
    │   │   ├── model           模型
    │   │   ├── repository      存储接口
    │   │   │   └── persistence ①存储接口实现   
    │   │   └── service         领域服务
    │   ├── interface           接口
    │   │   ├── handler         micro handler接口
    │   │   └── persistence     ②存储接口实现
    │   ├── registry            依赖注入，根据使用习惯，一般Go中不怎么喜欢这种方式
    │   └── usecase             应用用例
    │       ├── event           消息事件
    │       ├── service         应用服务
    ├── example         micro srv不同场景示例
    └── pb              基础服务协议统一.proto
```

### 系统架构图
<img src="/doc/img/architecture.png" width="50%">

### 业务架构图
*TODO*

**领域模型&整洁架构参考**
- [Clean Architecture in go](https://medium.com/@hatajoe/clean-architecture-in-go-4030f11ec1b1)
- [基于 DDD 的微服务设计和开发实战](https://www.infoq.cn/article/s_LFUlU6ZQODd030RbH9)
- [当中台遇上 DDD，我们该如何设计微服务？](https://www.infoq.cn/article/7QgXyp4Jh3-5Pk6LydWw)

## 快速开始

[Kubernetes环境](/deploy/k8s)

### 运行网关

自定义`micro`工具，[网关插件](/gateway/plugin.go)

#### 编译及运行

[gateway](/gateway)

#### Docker运行

*TODO*

### 运行服务
- Web应用
	- `app/console/web`控制台
- 聚合API
	- `app/console/api`控制台
- 基础服务
	- `srv/account`账户
	
```bash
$ cd {指定服务目录}

# 默认mdns注册中心
$ make build run

# 使用etcd注册中心
$ make build run registry=etcd
```

### Makefile
```bash
$ make build                                    # 编译
$ make run                                      # 运行
$ make run registry=etcd transport=tcp          # 运行，指定registry、transport
$ make build run                                # 编译&运行
$ make build run registry=etcd transport=tcp    # 编译&运行，指定registry、transport

$ make docker tag=xxx/xxx:v0.0.1
```
