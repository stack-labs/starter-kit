# Micro 快速开发工具包*项目进行中*

本仓库旨在提供面向Go-Micro生产环境的快速开发包。项目结合维护者们十余年的工作经验，不同领域的实战沉淀，一切为了缩短大家的选型、开发周期。

## 我们提供以下能力

- [ ] Web 服务
- [ ] 后台服务
- [ ] 网关
- [ ] 验证
- [ ] 链路追踪
- [ ] 配置中心
- [ ] 日志
- [ ] 服务代理
- [ ] ORM
- [ ] 单元测试
- [ ] 打包、部署
- [ ] Docker
- [ ] CICD
- [ ] K8s
- [ ] Service Mesh 服务网格
- [ ] GraphQL Gateway
- ...

![architecture](/doc/img/architecture.png "architecture")

*架构完善中*

## 快速开始

### 运行网关

自定义`micro`工具，[网关插件](/gateway)

#### 手动打包及运行

*略*

#### Docker运行

*TODO*

### 运行服务
- Web应用
	- `console/web`控制台
- 聚合API
	- `console/api`控制台
- 基础服务
	- `srv/account`账户
	
```bash
$ cd {指定服务目录}

# 默认mdns注册中心
$ make run

# 使用etcd注册中心
$ MICRO_REGISTRY=etcd make run
```

### 预制环境

环境安装

### 

待完善
