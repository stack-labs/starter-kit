# Micro 快速开发工具包*项目进行中*

本仓库旨在提供面向Go-Micro生产环境的快速开发包。项目结合维护者们十余年的工作经验，不同领域的实战沉淀，一切为了缩短大家的选型、开发周期。

## 目录

- 快速开始
    - [控制台示例](/console#目录)
    - [Hipster Shop示例](/hipstershop)
- [目标](#目标)
- [架构设计](#架构设计)
    - [目录结构](#目录结构)
- [部署](#部署)

## 目标

- 自定义[micro网关](gateway)
	- [x] `JWT`认证
	- [x] `Casbin`鉴权
	- [x] Tracing
	- [ ] RequestID
	- [x] Metrics
	- [ ] Access Log
	- ...
- 网关选择
	- micro api
		- [x] api
		- [x] rpc
		- proxy/http/web
		    - [x] [静态资源](/console/web/statik)
            - [x] [echo](/console/web/echo)
            - [x] [gin](/console/web/gin)
            - [x] [iris](/console/web/iris)
            - [x] [beego](/console/web/beego)
- 配置中心
- 前后端分离`console`
	- [x] [PanJiaChen/vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)
	- [ ] [tookit/vue-material-admin](https://github.com/tookit/vue-material-admin) 
	- [ ] [view-design/iview-admin](https://github.com/view-design/iview-admin)
- 参数验证
	- [x] [protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate)，适用于API`handler=rpc`的模式
	    - 规则配置[account.proto](/app/console/api/proto/account/account.proto#L21)
	    - 参数验证[account.go](/app/console/api/handler/account.go#L26)
- 领域驱动
	- [x] 整洁架构
- ORM
	- [x] gorm
	- [x] xorm
- 发布
	- [x] 灰度
	- [x] 蓝绿
	- *注:由于micro默认的api和web网关均不支持服务筛选，需要自己改造，方案参考[微服务协作开发、灰度发布之流量染色](https://micro.mu/blog/cn/2019/12/09/go-micro-service-chain.html)*
- 部署
	- [ ] K8S
		- [x] [helm](/deploy/k8s/helm)
	- [ ] Docker
- 安全
- CICD
	- [x] [Drone](https://drone.io/) [README](/deploy/docker/drone)
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

### 系统架构图
<img src="/doc/img/architecture.png" width="50%">

### 业务架构图

- [Console示例](/console)
- [Hipster Shop示例](/hipstershop)
    - 参考[GoogleCloudPlatform/microservices-demo](https://github.com/GoogleCloudPlatform/microservices-demo/)

**领域模型&整洁架构参考**
- [Clean Architecture in go](https://medium.com/@hatajoe/clean-architecture-in-go-4030f11ec1b1)
- [基于 DDD 的微服务设计和开发实战](https://www.infoq.cn/article/s_LFUlU6ZQODd030RbH9)
- [当中台遇上 DDD，我们该如何设计微服务？](https://www.infoq.cn/article/7QgXyp4Jh3-5Pk6LydWw)

## 部署

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

## 贡献

### 代码格式
- IDE IDEA/Goland，`Go->imports` 设置
    - Sorting type `gofmt`
    - [x] `Group stdlib imports`
        - [x] `Move all stdlib imports in a single group`
    - [x] `Move all imports in a single declaration`
