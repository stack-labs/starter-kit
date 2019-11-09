# Auth

集成`JWT`与`Casbin`
- [jwt-go](https://github.com/dgrijalva/jwt-go)
- [casbin](https://github.com/casbin/casbin)

## micro plugin 集成
- `auth`插件注册`casbin`的`adapter`
	- `auth.RegisterAdapter("default", a)`
	- 注册`key`与`--casbin_adapter`参数保持一致，默认`default`可以省略参数
- `--auth_pub_key`、`--casbin_model`均有默认路径，可以根据自己的目录进行配置
- `--casbin_watcher`可选
	- `auth.RegisterWatcher("default", w)`
	- 同样使用`default`可以省略参数
- 自定义`token`或`access control`验证失败时的响应
	- `auth.AuthResponse = auth.DefaultResponseHandler`
- `--casbin_public_user`为`public`接口指定`user`，默认为`public`，设置为空则无`public`接口

// micro plugin示例
```go
func init() {
	// adapter
	// xorm
	// a, _ := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/")
	// file
	a := fileadapter.NewAdapter("./conf/casbin_policy.csv")
	auth.RegisterAdapter("default", a)

	// watcher
	// https://casbin.org/docs/zh-CN/watchers
	// w, _ := rediswatcher.NewWatcher("127.0.0.1:6379")
	// auth.RegisterWatcher("default", w)

	// 自定义Response
	auth.AuthResponse = auth.DefaultResponseHandler

	api.Register(auth.NewPlugin())
}
```

```bash
--auth_pub_key value        Auth public key file (default: "./conf/auth_key.pub")
--casbin_model value        Casbin model config file (default: "./conf/casbin_model.conf")
--casbin_adapter value      Casbin registed adapter {default} (default: "default")
--casbin_watcher value      Casbin registed watcher {} (default: "default")
--casbin_public_user value  Casbin public user (default: "public")
```

