package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/micro/micro/api"
	"github.com/micro/micro/web"

	// micro plugins
	_ "github.com/micro/go-plugins/registry/kubernetes"
	_ "github.com/micro/go-plugins/transport/tcp"

	"github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/micro/go-micro/util/log"

	"github.com/micro-in-cn/starter-kit/gateway/plugin/auth"
	"github.com/micro-in-cn/starter-kit/gateway/plugin/cors"
	"github.com/micro-in-cn/starter-kit/gateway/plugin/metrics"
	"github.com/micro-in-cn/starter-kit/gateway/plugin/trace/opentracing"
	"github.com/micro-in-cn/starter-kit/gateway/plugin/util/response"
	tracer "github.com/micro-in-cn/starter-kit/pkg/opentracing"
)

var apiTracerCloser, webTracerCloser io.Closer

func pluginAfterFunc() error {
	// closer
	webTracerCloser.Close()
	apiTracerCloser.Close()

	return nil
}

func init() {
	// 跨域
	initCors()

	// 监控
	initMetrics()

	// Auth
	initAuth()

	// 链路追踪
	initTrace()
}

func initCors() {
	// 跨域
	corsPlugin := cors.NewPlugin(
		cors.WithAllowMethods(http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete),
		cors.WithAllowCredentials(true),
		cors.WithMaxAge(3600),
		cors.WithUseRsPkg(true),
	)
	api.Register(corsPlugin)
	web.Register(corsPlugin)
}

func initAuth() {
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

	authPlugin := auth.NewPlugin(
		auth.WithResponseHandler(response.DefaultResponseHandler),
		auth.WithSkipperFunc(func(r *http.Request) bool {
			return false
		}),
	)
	api.Register(authPlugin)
	// web.Register(authPlugin)
}

func initMetrics() {
	api.Register(metrics.NewPlugin(
		metrics.WithSkipperFunc(func(r *http.Request) bool {
			return false
		}),
	))

	web.Register(metrics.NewPlugin(
		metrics.WithSkipperFunc(func(r *http.Request) bool {
			path := r.URL.Path
			idx := strings.Index(path[1:], "/")
			if idx > 0 {
				path = path[idx+1:]
			}
			if strings.HasPrefix(path, "/v1/") {
				return false
			}
			return true
		}),
	))
}

func initTrace() {
	apiTracer, apiCloser, err := tracer.NewJaegerTracer("go.micro.api", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	apiTracerCloser = apiCloser
	api.Register(opentracing.NewPlugin(
		opentracing.WithTracer(apiTracer),
	))

	webTracer, webCloser, err := tracer.NewJaegerTracer("go.micro.web", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	webTracerCloser = webCloser
	web.Register(opentracing.NewPlugin(
		opentracing.WithTracer(webTracer),
		opentracing.WithSkipperFunc(func(r *http.Request) bool {
			path := r.URL.Path
			idx := strings.Index(path[1:], "/")
			if idx > 0 {
				path = path[idx+1:]
			}
			if strings.HasPrefix(path, "/v1/") {
				return false
			}
			return true
		}),
	))
}
