package main

import (
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/pkg/cli"
	"github.com/stack-labs/stack-rpc/pkg/config"
	"github.com/stack-labs/stack-rpc/pkg/config/source/file"
	"github.com/stack-labs/stack-rpc/util/log"

	"github.com/stack-labs/starter-kit/console/account/conf"
	"github.com/stack-labs/starter-kit/console/account/interface/handler"
	"github.com/stack-labs/starter-kit/console/account/registry"
	"github.com/stack-labs/starter-kit/pkg/plugin/wrapper/select/chain"
	"github.com/stack-labs/starter-kit/pkg/plugin/wrapper/trace/opentracing"
	"github.com/stack-labs/starter-kit/pkg/tracer"
)

func main() {
	md := make(map[string]string)
	md["chain"] = "gray"

	// New Service
	service := stack.NewService(
		stack.Name("stack.rpc.srv.account"),
		stack.Version("latest"),
		stack.Metadata(md),
		stack.Flags(
			cli.StringFlag{
				Name:  "conf_path",
				Value: "./conf/",
				Usage: "配置文件目录",
			},
		),
		stack.Action(func(ctx *cli.Context) {
			confPath := ctx.String("conf_path")
			conf.BASE_PATH = confPath

			// 配置加载
			cfg, _ := config.NewConfig()
			err := cfg.Load(file.NewSource(
				file.WithPath(conf.BASE_PATH + "config.yaml"),
			))
			if err != nil {
				log.Fatal(err)
			}
		}),
	)

	// 链路追踪
	t, closer, err := tracer.NewJaegerTracer("stack.rpc.srv.account", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()
	service.Init(
		// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
		// TODO stack-rpc
		stack.WrapCall(opentracing.NewCallWrapper(t)),
		stack.WrapHandler(opentracing.NewHandlerWrapper(t)),
		stack.WrapClient(chain.NewClientWrapper()),
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
