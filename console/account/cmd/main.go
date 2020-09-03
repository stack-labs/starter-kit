package main

import (
	"context"
	"os"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v3/config"
	"github.com/micro/go-micro/v3/config/source/file"
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/profile"
	"github.com/micro/micro/v3/service"

	"github.com/micro-in-cn/starter-kit/console/account/conf"
	"github.com/micro-in-cn/starter-kit/console/account/interface/handler"
	"github.com/micro-in-cn/starter-kit/console/account/registry"
	tracer "github.com/micro-in-cn/starter-kit/pkg/opentracing"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/trace/opentracing"
	_ "github.com/micro-in-cn/starter-kit/profile"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "profile",
			Usage: "micro profile",
		},
		// TODO V3 命令行参数报错
		&cli.StringFlag{
			Name:  "conf_path",
			Value: "./conf/",
			Usage: "配置文件目录",
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
		confPath := ctx.String("conf_path")
		conf.BASE_PATH = confPath

		// TODO 配置加载，不在开箱即用
		_, err := config.NewConfig(
			config.WithSource(
				file.NewSource(
					file.WithPath(conf.BASE_PATH + "config.yaml"),
				),
			),
		)
		if err != nil {
			return err
		}

		return run()
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

func run() error {
	md := make(map[string]string)
	md["chain"] = "gray"

	// New Service
	svc := service.New(
		service.Name("go.micro.srv.account"),
		service.Version("latest"),
		service.Metadata(md),
	)

	// 链路追踪
	t, closer, err := tracer.NewJaegerTracer("go.micro.srv.account", "127.0.0.1:6831")
	if err != nil {
		logger.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()

	svc.Init(
		// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
		service.WrapCall(opentracing.NewCallWrapper(t)),
		service.WrapHandler(opentracing.NewHandlerWrapper(t)),
		//service.WrapClient(chain.NewClientWrapper()),
	)

	// Initialise service
	svc.Init()

	c, err := registry.NewContainer()
	if err != nil {
		logger.Fatalf("failed to build container: %v", err)
	}

	// Register Handler
	handler.Apply(svc.Server(), c)

	// Run service
	return svc.Run()
}
