package main

//go:generate statik -src=./vue/dist -dest=./ -f
import (
	"github.com/micro-in-cn/x-gateway/plugin/opentracing"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"

	"github.com/micro-in-cn/starter-kit/app/console/web/beego"
	"github.com/micro-in-cn/starter-kit/app/console/web/echo"
	"github.com/micro-in-cn/starter-kit/app/console/web/gin"
	"github.com/micro-in-cn/starter-kit/app/console/web/iris"
	"github.com/micro-in-cn/starter-kit/app/console/web/statik"
	tracer "github.com/micro-in-cn/starter-kit/pkg/opentracing"
)

func main() {
	md := make(map[string]string)
	md["chain"] = "gray"

	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.console"),
		web.Version("latest"),
		web.Metadata(md),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// 链路追踪
	t, closer, err := tracer.NewJaegerTracer("go.micro.web.console", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()

	// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
	h := opentracing.NewPlugin(
		opentracing.WithTracer(t),
		opentracing.WithAutoStart(false),
	).Handler()

	// TODO Path末尾"/"问题
	// Echo
	echoHandler, err := echo.New()
	if err != nil {
		log.Fatal(err)
	}
	service.Handle("/v1/echo/", h(echoHandler))

	// Gin
	ginHandler, err := gin.New()
	if err != nil {
		log.Fatal(err)
	}
	service.Handle("/v1/gin/", h(ginHandler))

	// Iris
	irisHandler, err := iris.New()
	if err != nil {
		log.Fatal(err)
	}
	service.Handle("/v1/iris/", h(irisHandler))

	// Beego
	beegoHandler, err := beego.New()
	if err != nil {
		log.Fatal(err)
	}
	service.Handle("/v1/beego/", h(beegoHandler))

	// register static file handler
	// 使用statik打包需要:make statik，编译时增加`-tags "statik"`标签
	service.Handle("/", statik.Handler())

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
