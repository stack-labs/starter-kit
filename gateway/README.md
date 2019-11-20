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
$ make run_api                                  # 默认mdns + http
$ make run_api registry=etcd transport=tcp      # 指定etcd + tcp

# Web
$ make run_web                                  # 默认mdns + http
$ make run_web registry=etcd transport=tcp      # 指定etcd + tcp
```

## Docker

```bash
# tag自定义
$ make docker tag=hbchen/micro:v1.14.0_starter_kit
```
