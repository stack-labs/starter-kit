package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v3/api"
	"github.com/micro/go-micro/v3/logger"
	log "github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/profile"

	"github.com/hb-go/micro-plugins/v3/web"
	"github.com/micro-in-cn/starter-kit/console/web/beego"
	"github.com/micro-in-cn/starter-kit/console/web/echo"
	"github.com/micro-in-cn/starter-kit/console/web/gin"
	"github.com/micro-in-cn/starter-kit/console/web/iris"
	"github.com/micro-in-cn/starter-kit/console/web/statik"
	tracer "github.com/micro-in-cn/starter-kit/pkg/opentracing"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/opentracing"
	_ "github.com/micro-in-cn/starter-kit/profile"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "profile",
			Usage: "micro profile",
		},
		&cli.BoolFlag{
			Name:  "debug",
			Usage: "starter kit debug.",
		},
	}

	app.Before = func(ctx *cli.Context) error {
		p := ctx.String("profile")
		// apply the profile
		if profile, err := profile.Load(p); err != nil {
			logger.Fatal(err)
		} else {
			// load the profile
			profile.Setup(ctx)
		}

		return nil
	}

	app.Action = func(ctx *cli.Context) error {
		return run(ctx)
	}

	app.Commands = cli.Commands{
		&cli.Command{
			Name:  "reload",
			Usage: "TODO",
			Action: func(ctx *cli.Context) error {
				return nil
			},
		},
	}

	ctx := context.TODO()
	if err := app.RunContext(ctx, os.Args); err != nil {
		logger.Fatal(err)
	}
}

func run(ctx *cli.Context) error {
	md := make(map[string]string)
	md["chain"] = "gray"

	// create new web service
	service := web.NewService(
		web.Name("go.micro.api.consoleweb"),
		web.Version("latest"),
		web.Metadata(md),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// 链路追踪
	t, closer, err := tracer.NewJaegerTracer("go.micro.api.console.web", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()

	// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
	traceHandler := opentracing.NewPlugin(
		opentracing.WithTracer(t),
		opentracing.WithAutoStart(false),
		opentracing.WithSkipperFunc(func(r *http.Request) bool {
			return false
		}),
	).Handler()

	mux := http.NewServeMux()

	// TODO Path末尾"/"问题
	// Echo
	echoHandler, err := echo.New()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/v1/echo/", traceHandler(echoHandler))

	// Gin
	ginHandler, err := gin.New()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/v1/gin/", traceHandler(ginHandler))

	// Iris
	irisHandler, err := iris.New()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/v1/iris/", traceHandler(irisHandler))

	// Beego
	beegoHandler, err := beego.New()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/v1/beego/", traceHandler(beegoHandler))

	// register static file handler
	// 使用statik打包需要:make statik，编译时增加`-tags "statik"`标签
	mux.Handle("/", statik.Handler())

	// Path前缀
	h := &handler{
		prefix: "/console",
		mux:    mux,
	}
	service.Handle("/console/", h, &api.Endpoint{
		Name:    "console",
		Path:    []string{"^/console/*"},
		Method:  []string{"POST", "GET", "DELETE", "HEAD", "OPTIONS"},
		Handler: "http",
	})

	// run service
	return service.Run()
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
