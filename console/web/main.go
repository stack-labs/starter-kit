package main

//go:generate statik -src=./vue/dist -dest=./ -f
import (
	"net/http"
	"strings"

	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/api"
	"github.com/stack-labs/stack-rpc/service/web"
	"github.com/stack-labs/stack-rpc/util/log"
	"github.com/stack-labs/starter-kit/console/web/beego"
	"github.com/stack-labs/starter-kit/console/web/echo"
	"github.com/stack-labs/starter-kit/console/web/gin"
	"github.com/stack-labs/starter-kit/console/web/iris"
	"github.com/stack-labs/starter-kit/console/web/statik"
	//"github.com/stack-labs/starter-kit/pkg/plugin/opentracing"
	//"github.com/stack-labs/starter-kit/pkg/tracer"
)

func main() {
	md := make(map[string]string)
	md["chain"] = "gray"

	// 链路追踪
	//t, closer, err := tracer.NewJaegerTracer("stack.rpc.api.console.web", "127.0.0.1:6831")
	//if err != nil {
	//	log.Fatalf("opentracing tracer create error:%v", err)
	//}
	//defer closer.Close()

	// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
	//traceHandler := opentracing.Handler(
	//	opentracing.WithTracer(t),
	//	opentracing.WithAutoStart(false),
	//	opentracing.WithSkipperFunc(func(r *http.Request) bool {
	//		return false
	//	}),
	//)

	mux := http.NewServeMux()

	// Path前缀
	h := &handler{
		prefix: "/console",
		mux:    mux,
	}

	// create new web service
	service := stack.NewWebService(
		stack.Name("stack.rpc.api.web"),
		stack.Version("latest"),
		stack.Metadata(md),
		web.RootPath("/console"),
		web.HandleFuncs(web.HandlerFunc{
			Route: "/",
			Func:  h.ServeHTTP,
		}),
		web.HandlerOptions(api.WithEndpoint(&api.Endpoint{
			Name:    "console",
			Path:    []string{"^/console/*"},
			Method:  []string{"POST", "GET", "DELETE", "HEAD", "OPTIONS"},
			Handler: "proxy",
		})),
	)

	service.Init()

	// TODO Path末尾"/"问题
	// Echo
	echoHandler, err := echo.New()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/v1/echo/", echoHandler)

	// Gin
	ginHandler, err := gin.New()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/v1/gin/", ginHandler)

	// Iris
	irisHandler, err := iris.New()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/v1/iris/", irisHandler)

	// Beego
	beegoHandler, err := beego.New()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/v1/beego/", beegoHandler)

	// register static file handler
	// 使用statik打包需要:make statik，编译时增加`-tags "statik"`标签
	mux.Handle("/", statik.Handler())

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type handler struct {
	prefix string
	mux    *http.ServeMux
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debugf("console receive request path: %v", r.URL.Path)
	r.URL.Path = strings.TrimPrefix(r.URL.Path, h.prefix)
	h.mux.ServeHTTP(w, r)
}
