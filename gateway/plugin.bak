// +build ignore

package main

import (
	"io"
	"net/http"
	"time"

	"github.com/stack-labs/stack-rpc-plugins/service/stackway/plugin"
	"github.com/stack-labs/stack-rpc/util/log"
	"golang.org/x/time/rate"

	"github.com/stack-labs/starter-kit/pkg/plugin/auth"
	"github.com/stack-labs/starter-kit/pkg/plugin/chain"
	"github.com/stack-labs/starter-kit/pkg/plugin/cors"
	"github.com/stack-labs/starter-kit/pkg/plugin/metrics"
	"github.com/stack-labs/starter-kit/pkg/plugin/opentracing"
	"github.com/stack-labs/starter-kit/pkg/tracer"
	"github.com/stack-labs/starter-kit/pkg/utils/response"
)

var apiTracerCloser io.Closer

func pluginAfterFunc() error {
	// closer
	return apiTracerCloser.Close()
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
	if err := plugin.Register(corsPlugin); err != nil {
		log.Error(err)
	}
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

	if err := plugin.Register(authPlugin); err != nil {
		log.Error(err)
	}
}

func initMetrics() {
	metricsPlugin := metrics.NewPlugin(
		metrics.WithNamespace("gateway"),
		metrics.WithSubsystem(""),
		metrics.WithSkipperFunc(func(r *http.Request) bool {
			return false
		}),
	)

	if err := plugin.Register(metricsPlugin); err != nil {
		log.Error(err)
	}
}

// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
func initTrace() {
	apiTracer, apiCloser, err := tracer.NewJaegerTracer("stack.rpc.api", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}

	limiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 10)
	apiTracerCloser = apiCloser

	otPlugin := opentracing.NewPlugin(
		opentracing.WithTracer(apiTracer),
		opentracing.WithSkipperFunc(func(r *http.Request) bool {
			// 采样频率控制，根据需要细分Host、Path等分别控制
			return !limiter.Allow()
		}),
	)

	if err := plugin.Register(otPlugin); err != nil {
		log.Error(err)
	}
}

func initChain() {
	chainPlugin := chain.New(chain.WithChainsFunc(func(r *http.Request) []string {
		return []string{"gray"}
	}))

	if err := plugin.Register(chainPlugin); err != nil {
		log.Error(err)
	}
}
