# Gateway

自定义micro工具，增加`Auth`、`CORS`插件

- 认证&鉴权`JWT`+`Casbin` [Auth](https://github.com/hb-go/micro-plugins/tree/master/micro/auth)
- 跨域支持 [CORS](https://github.com/hb-go/micro-plugins/tree/master/micro/cors)

## 运行网关

```bash
# 打包
$ make build

# 运行
$ ./micro --registry=etcd api
$ ./micro --registry=etcd web
```
