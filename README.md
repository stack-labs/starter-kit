# Micro 快速开发工具包

> *项目进行中*

本仓库旨在提供面向Go-Micro生产环境的快速开发包。

## 目录

- 快速开始示例
    - [控制台示例](/console#目录)
        - 以最常见的登录流程为例，实现一个场景简单，但包含微服务各种治理能力的示例
    - [Hipster Shop示例](/hipstershop)
        - 参考[GoogleCloudPlatform/microservices-demo](https://github.com/GoogleCloudPlatform/microservices-demo/)实现一个业务场景比较复杂的微服务应用
- [架构设计](#架构设计)
- [目录结构](#目录结构)
- [目标功能](#目标功能)
- [开发环境](#开发环境)
- [部署环境](#部署环境)
- [参与贡献](#参与贡献)

## 架构设计

### 系统架构图
<img src="/doc/img/architecture.png" width="50%">

### 业务架构图

**[Console示例](/console)**

<img src="/doc/img/console-design.png" width="50%">

- [Hipster Shop示例](/hipstershop)
    - 参考[GoogleCloudPlatform/microservices-demo](https://github.com/GoogleCloudPlatform/microservices-demo/)

**领域模型&整洁架构参考**
- [Clean Architecture in go](https://medium.com/@hatajoe/clean-architecture-in-go-4030f11ec1b1)
- [基于 DDD 的微服务设计和开发实战](https://www.infoq.cn/article/s_LFUlU6ZQODd030RbH9)
- [当中台遇上 DDD，我们该如何设计微服务？](https://www.infoq.cn/article/7QgXyp4Jh3-5Pk6LydWw)

## 目录结构

```bash
├── console             控制台示例
│   ├── account         go.micro.srv.account，Account服务
│   │   ├── domain              领域
│   │   │   ├── model           模型
│   │   │   ├── repository      存储接口
│   │   │   │   └── persistence ①存储接口实现   
│   │   │   └── service         领域服务
│   │   ├── interface           接口
│   │   │   ├── handler         micro handler接口
│   │   │   └── persistence     ②存储接口实现
│   │   ├── registry            依赖注入，根据使用习惯，一般Go中不怎么喜欢这种方式
│   │   └── usecase             应用用例
│   │       ├── event           消息事件
│   │       └── service         应用服务
│   ├── api             go.micro.api.console，API服务
│   ├── pb              服务协议统一.proto
│   └── web             go.micro.api.console，Web服务，集成gin、echo、iris等web框架
├── deploy              部署
│   ├── docker
│   └── k8s
├── doc                 文档资源
├── gateway             网关，自定义micro
└── pkg                 公共资源包
```

## 目标功能

- 自定义[micro网关](gateway)
	- [x] `JWT`认证
	- [x] `Casbin`鉴权
	- [x] Tracing
	- [ ] RequestID
	- [x] Metrics
	- [ ] Access Log
	- ...
- API服务
    - 网关使用默认处理器(`handler=meta`)，聚合服务通过`Endpoint`定义路由规则，实现统一网关管理`rpc`和`http`类型的聚合服务
        - *注:`go-micro/web`服务注册不支持`Endpoint`定义，需要自定义`web.Service`([实现参考](https://github.com/hb-go/micro-plugins/tree/master/web))，[issue#1097](https://github.com/micro/go-micro/issues/1097)*
	- [x] api
    - [x] rpc
    - proxy/http/web
        - [x] [静态资源](/console/web/statik)
            - *前后端分离场景将静态资源独立更好，但不排除使用Web模板框架的应用加入微服务体系，尤其在已有单体逐步拆分的演进过程中*
        - [x] [echo](/console/web/echo)
        - [x] [gin](/console/web/gin)
        - [x] [iris](/console/web/iris)
        - [x] [beego](/console/web/beego)
    - API文档
        - [x] swagger
            - *使用[grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)的`protoc-gen-swagger`生成swagger文档，适用于API`handler=rpc`的模式*
            - [示例](/console#Swagger文档生成)
- 配置中心
    - [ ] XConf
- 前后端分离`console`
	- [x] [PanJiaChen/vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)
	    - [示例](/console/web/vue)
	- [ ] [tookit/vue-material-admin](https://github.com/tookit/vue-material-admin) 
	- [ ] [view-design/iview-admin](https://github.com/view-design/iview-admin)
- 参数验证
	- [x] [protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate)，适用于API`handler=rpc`的模式
	    - 规则配置[account.proto](/console/pb/api/account.proto#L21)
	    - 参数验证[account.go](/console/api/handler/account.go#L26)
- 领域驱动
	- [x] 整洁架构
	    - 示例[console/account](/console/account)
- ORM
	- [x] gorm
	    - [示例](/console/account/domain/repository/persistence/gorm)
	- [x] xorm
	    - [示例](/console/account/domain/repository/persistence/xorm)
- 发布
	- [x] 灰度
	- [x] 蓝绿
	- *注:由于micro默认的api和web网关均不支持**服务筛选**，需要自己改造，方案参考[微服务协作开发、灰度发布之流量染色](https://micro.mu/blog/cn/2019/12/09/go-micro-service-chain.html)*
- 部署
	- K8S
		- [x] [helm](/deploy/k8s/helm)
	- [x] Docker
	    - 示例[console](/console/docker-compose.yml)
- 安全
- CICD
	- [Drone](https://drone.io/) [README](/deploy/docker/drone)
	    - [x] Go & Node编译
	    - [x] Docker镜像
	    - [x] Kubernetes发布
	    - [x] 缓存
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

## 开发环境

- 本地
    - [x] [Docker Compose](/console#docker-compose启动)
- 在线
    - [x] CICD
        - [x] Drone
        - [x] Kubernetes
    - [x] 本地服务接入
        - [x] Docker环境
        - [ ] K8S环境
        
### 本地服务接入-Network代理

<details>
  <summary> 本地服务接入-Network代理 </summary>

**Network代理测试**

以`console`的[docker-compose.yaml](/console/docker-compose.yml)为例，假设`compose`为在线环境，本地开发`account`服务。
- `compose`中加入`network`服务，参考[docker-compose-network.yml](/console/docker-compose-network.yml)
- `api`服务使用`network`做代理`MICRO_PROXY=go.micro.network`，***注意测试时`compose`中仅`api`服务使用代理***
- 本地启动`network`
    - `micro --registry=etcd --transport=tcp network --nodes=127.0.0.1:8085 --address=:8086 --advertise_strategy=local`
- 剩下的工作使用`proxy`对`route`的筛选功能，参考PR[#897](https://github.com/micro/go-micro/pull/897)
    - 查看本地`router`，`micro --registry=etcd --transport=tcp network routes`
    - `curl -XPOST -H "Micro-Router: 6832f8ff-1217-4119-8a56-9a90adf19fef" -d '{"username" : "admin","password":"123456"}' http://localhost:8080/account/login`
    
**真实场景**

- 考虑所有服务都可以自助路由到本地，不能直接使用`Micro-Router`(因为`Micro-Router`会在全链路生效)，可以自定义`header`来定义`router`筛选的应用范围，通过`Client/Call Wrap`实现，参考实现[router_filter](/pkg/plugin/wrapper/client/router_filter)
- 要做到`api`服务可以路由筛选，在网关层与[流量染色](https://micro.mu/blog/cn/2019/12/09/go-micro-service-chain.html)有相同的问题，不支持**服务筛选**，需要去掉`SelectOption`
- 网关及服务全部使用`network`做代理

```shell script
# api服务路由到本地
curl -XPOST -H "X-Micro-Router-Filter: go.micro.api.console:6832f8ff-1217-4119-8a56-9a90adf19fef" -d '{"username" : "admin","password":"123456"}' http://localhost:8080/account/login

# account服务理由到本地
curl -XPOST -H "X-Micro-Router-Filter: go.micro.srv.account:6832f8ff-1217-4119-8a56-9a90adf19fef" -d '{"username" : "admin","password":"123456"}' http://localhost:8080/account/login

# api和account服务都路由到本地
curl -XPOST -H "X-Micro-Router-Filter: go.micro.api.console:6832f8ff-1217-4119-8a56-9a90adf19fef;go.micro.srv.account:6832f8ff-1217-4119-8a56-9a90adf19fef" -d '{"username" : "admin","password":"123456"}' http://localhost:8080/account/login
```
  
</details>

## 生产环境

[Kubernetes环境](/deploy/k8s)

## 可选服务

<details>
  <summary> Jaeger </summary>

> 浏览器访问:http://localhost:16686/
```bash
$ docker run -d --name=jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp   -p5778:5778 -p16686:16686 -p14268:14268 -p9411:9411 jaegertracing/all-in-one:latest
```

</details>

<details>
  <summary> Prometheus </summary>

> 浏览器访问:http://localhost:9090/

> `prometheus.yml`参考`gateway`插件`[metrics/prometheus.yml](/gateway/plugin/metrics/prometheus.yml)
```bash
$ docker run -d --name prometheus -p 9090:9090 -v ~/tmp/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
```

</details>

<details>
  <summary> Grafana </summary>

> 浏览器访问:http://localhost:3000/

> `Grafana`仪表盘`import`[metrics/grafan.json](/gateway/plugin/metrics/grafan.json)
```bash
$ docker run --name grafana -d -p 3000:3000 grafana/grafana
```

</details>

## 参与贡献

### 代码格式
- IDE IDEA/Goland，`Go->imports` 设置
    - Sorting type `gofmt`
    - [x] `Group stdlib imports`
        - [x] `Move all stdlib imports in a single group`
    - [x] `Move all imports in a single declaration`
