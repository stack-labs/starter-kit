package main

import (
	"context"
	"os"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/profile"
	"github.com/micro/micro/v3/service"

	"github.com/micro-in-cn/starter-kit/console/api/client"
	"github.com/micro-in-cn/starter-kit/console/api/handler"
	tracer "github.com/micro-in-cn/starter-kit/pkg/opentracing"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/client/router_filter"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/trace/opentracing"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/validate"
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
		service.Name("go.micro.api.console"),
		service.Version("v1"),
		service.Metadata(md),
	)

	// 链路追踪
	t, closer, err := tracer.NewJaegerTracer("go.micro.api.console", "127.0.0.1:6831")
	if err != nil {
		logger.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()

	// Initialise service
	svc.Init(
		// 流量染色
		//service.WrapClient(chain.NewClientWrapper()),
		// 路由筛选
		service.WrapCall(router_filter.NewCallWrapper()),
	)
	svc.Init(
		// create wrap for the Example srv client
		service.WrapHandler(client.AccountWrapper(svc)),
		service.WrapHandler(validate.NewHandlerWrapper()),
		service.WrapCall(validate.NewCallWrapper()),
		// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
		service.WrapHandler(opentracing.NewHandlerWrapper(t)),         // server端handler接受请求
		service.WrapSubscriber(opentracing.NewSubscriberWrapper(nil)), // server端subscriber接受消息
		service.WrapClient(opentracing.NewClientWrapper(nil)),         // client端发起请求，包括Call()、Stream()、Publish()
		service.WrapCall(opentracing.NewCallWrapper(t)),               // client端发起请求，仅Call()
	)

	// Register Handler
	handler.RegisterHandler(svc.Server())

	// Run service
	return svc.Run()
}
