package main

import (
	"net/http"

	"github.com/micro/micro/api"
	"github.com/micro/micro/web"

	// micro plugins
	_ "github.com/micro/go-plugins/registry/kubernetes"
	_ "github.com/micro/go-plugins/transport/tcp"

	"github.com/casbin/casbin/v2/persist/file-adapter"

	"github.com/micro-in-cn/starter-kit/gateway/plugin/auth"
	"github.com/micro-in-cn/starter-kit/gateway/plugin/cors"
	"github.com/micro-in-cn/starter-kit/gateway/plugin/metrics"
)

// API
func init() {
	// 跨域
	corsPlugin := cors.NewPlugin(
		cors.WithAllowMethods(http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete),
		cors.WithAllowCredentials(true),
		cors.WithMaxAge(3600),
		cors.WithUseRsPkg(true),
	)
	api.Register(corsPlugin)
	web.Register(corsPlugin)

	// adapter
	// xorm
	// a, _ := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/")
	// file
	a := fileadapter.NewAdapter("./conf/casbin_policy.csv")
	auth.RegisterAdapter("default", a)

	// watcher
	// https://casbin.org/docs/zh-CN/watchers
	// w := etcdwatcher.NewWatcher("http://127.0.0.1:2379")
	// w, _ := rediswatcher.NewWatcher("127.0.0.1:6379")
	// auth.RegisterWatcher("default", w)

	// 自定义Response
	auth.AuthResponse = auth.DefaultResponseHandler

	authPlugin := auth.NewPlugin(
		auth.WithSkipperFunc(func(r *http.Request) bool {
			return false
		}),
	)
	api.Register(authPlugin)
	// web.Register(authPlugin)

	metricsPlugin := metrics.NewPlugin(
		metrics.WithSkipperFunc(func(r *http.Request) bool {
			return false
		}),
	)
	api.Register(metricsPlugin)
	// web.Register(metricsPlugin)
}
