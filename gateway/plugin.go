// +build api

package main

import (
	"net/http"

	// tcp transport
	_ "github.com/micro/go-plugins/transport/tcp"
	// k8s registry
	_ "github.com/micro/go-plugins/registry/kubernetes"

	"github.com/hb-go/micro-plugins/micro/auth"
	"github.com/hb-go/micro-plugins/micro/cors"
	"github.com/micro/micro/api"
)

func init() {
	// 跨域
	api.Register(cors.NewPlugin(
		cors.WithAllowMethods(http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete),
		cors.WithAllowCredentials(true),
		cors.WithMaxAge(3600),
		cors.WithUseRsPkg(true),
	))

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
