package main

import (
	"io"
	"net/http"
	"time"

	"github.com/casbin/casbin/v2/persist/file-adapter"
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"
	_ "github.com/micro/go-plugins/transport/tcp/v2"
	"github.com/micro/micro/v3/plugin"
	"github.com/micro/micro/v3/service/logger"
	"golang.org/x/time/rate"

	tracer "github.com/micro-in-cn/starter-kit/pkg/opentracing"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/auth"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/cors"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/metrics"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/micro/chain"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/opentracing"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/utils/response"
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
	plugin.Register(corsPlugin, plugin.Module("gateway"))
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
	if err := plugin.Register(authPlugin, plugin.Module("gateway")); err != nil {
		logger.Error(err)
	}
}

func initMetrics() {
	plugin.Register(metrics.NewPlugin(
		metrics.WithNamespace("gateway"),
		metrics.WithSubsystem(""),
		metrics.WithSkipperFunc(func(r *http.Request) bool {
			return false
		}),
	), plugin.Module("gateway"))
}

// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
func initTrace() {
	apiTracer, apiCloser, err := tracer.NewJaegerTracer("go.micro.api", "127.0.0.1:6831")
	if err != nil {
		logger.Fatalf("opentracing tracer create error:%v", err)
	}

	limiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 10)
	apiTracerCloser = apiCloser
	plugin.Register(opentracing.NewPlugin(
		opentracing.WithTracer(apiTracer),
		opentracing.WithSkipperFunc(func(r *http.Request) bool {
			// 采样频率控制，根据需要细分Host、Path等分别控制
			if !limiter.Allow() {
				return true
			}
			return false
		}),
	), plugin.Module("gateway"))
}

func initChain() {
	// 在网关创建染色条件，将覆盖客户端
	plugin.Register(chain.New(chain.WithChainsFunc(func(r *http.Request) []string {
		return []string{"gray"}
	})), plugin.Module("gateway"))
}
