# Gateway

自定义micro工具，增加`Auth`、`CORS`等插件

- 认证&鉴权`JWT`+`Casbin` [Auth](/gateway/plugin/auth)
- 跨域支持 [CORS](/gateway/plugin/cors)
- Metrics [Prometheus](/gateway/plugin/metrics)

## 运行网关

```bash
# 编译
$ make build

# API
$ make run_api
$ make run_api registry=etcd transport=tcp    # 指定registry

# Web
$ make run_web
$ make run_web registry=etcd transport=tcp    # 指定registry
```

## Docker

```bash
# tag自定义
$ make docker tag=hbchen/micro:v1.14.0_starter_kit
```
