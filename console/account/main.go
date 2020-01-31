package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"

	"github.com/micro-in-cn/starter-kit/console/account/conf"
	"github.com/micro-in-cn/starter-kit/console/account/interface/handler"
	"github.com/micro-in-cn/starter-kit/console/account/registry"
	tracer "github.com/micro-in-cn/starter-kit/pkg/opentracing"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/select/chain"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/trace/opentracing"
)

func main() {
	md := make(map[string]string)
	md["chain"] = "gray"

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.account"),
		micro.Version("latest"),
		micro.Metadata(md),
		micro.Flags(
			&cli.StringFlag{
				Name:  "conf_path",
				Value: "./conf/",
				Usage: "配置文件目录",
			},
		),
		micro.Action(func(ctx *cli.Context) error {
			confPath := ctx.String("conf_path")
			conf.BASE_PATH = confPath

			// 配置加载
			err := config.LoadFile(conf.BASE_PATH + "config.yaml")

			return err
		}),
	)

	// 链路追踪
	t, closer, err := tracer.NewJaegerTracer("go.micro.srv.account", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()
	service.Init(
		// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
		micro.WrapCall(opentracing.NewCallWrapper(t)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(t)),
		micro.WrapClient(chain.NewClientWrapper()),
	)

	// Initialise service
	service.Init()

	ctn, err := registry.NewContainer()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	// Register Handler
	handler.Apply(service.Server(), ctn)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
