package main

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/micro/micro/v2/api"
	"github.com/micro/micro/v2/web"
	"golang.org/x/time/rate"

	// micro plugins
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"
	_ "github.com/micro/go-plugins/transport/tcp/v2"

	"github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/micro/go-micro/v2/util/log"

	"github.com/micro-in-cn/starter-kit/pkg/plugin/micro/chain"
	"github.com/micro-in-cn/x-gateway/plugin/auth"
	"github.com/micro-in-cn/x-gateway/plugin/cors"
	"github.com/micro-in-cn/x-gateway/plugin/metrics"
	"github.com/micro-in-cn/x-gateway/plugin/opentracing"
	tracer "github.com/micro-in-cn/x-gateway/plugin/trace"
	"github.com/micro-in-cn/x-gateway/utils/response"
)

var apiTracerCloser, webTracerCloser io.Closer

func pluginAfterFunc() error {
	// closer
	webTracerCloser.Close()
	apiTracerCloser.Close()

	return nil
}

// 插件注册
func init() {
	// 跨域
	initCors()

	// 监控
	initMetrics()

	// Auth
	initAuth()

	// 链路追踪
	initTrace()

	// 流量染色
	initChain()

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

	webAuthPlugin := auth.NewPlugin(
		auth.WithResponseHandler(response.DefaultResponseHandler),
		auth.WithSkipperFunc(func(r *http.Request) bool {
			// 自定义skipper规则
			return true
		}),
	)
	web.Register(webAuthPlugin)
}

func initMetrics() {
	api.Register(metrics.NewPlugin(
		metrics.WithNamespace("gateway"),
		metrics.WithSubsystem(""),
		metrics.WithSkipperFunc(func(r *http.Request) bool {
			return false
		}),
	))

	web.Register(metrics.NewPlugin(
		metrics.WithNamespace("gateway"),
		metrics.WithSubsystem(""),
		metrics.WithSkipperFunc(func(r *http.Request) bool {
			// 过滤micro web服务的前缀，便于设置统一规则，如/console/v1/* => /v1/*
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

// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
func initTrace() {
	apiTracer, apiCloser, err := tracer.NewJaegerTracer("go.micro.api", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}

	limiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 10)
	apiTracerCloser = apiCloser
	api.Register(opentracing.NewPlugin(
		opentracing.WithTracer(apiTracer),
		opentracing.WithSkipperFunc(func(r *http.Request) bool {
			// 采样频率控制，根据需要细分Host、Path等分别控制
			if !limiter.Allow() {
				return true
			}
			return false
		}),
	))

	webTracer, webCloser, err := tracer.NewJaegerTracer("go.micro.web", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	webTracerCloser = webCloser
	web.Register(opentracing.NewPlugin(
		opentracing.WithTracer(webTracer),
		opentracing.WithSkipperFunc(func(r *http.Request) bool {
			// Host、Path等过滤规则
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

func initChain() {
	api.Register(chain.New(chain.WithChainsFunc(func(r *http.Request) []string {
		return []string{"gray"}
	})))

	web.Register(chain.New(chain.WithChainsFunc(func(r *http.Request) []string {
		return []string{"gray"}
	})))
}
