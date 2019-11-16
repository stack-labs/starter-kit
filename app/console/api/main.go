package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	"github.com/micro-in-cn/starter-kit/app/console/api/client"
	"github.com/micro-in-cn/starter-kit/app/console/api/handler"
	tracer "github.com/micro-in-cn/starter-kit/pkg/opentracing"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/trace/opentracing"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.console"),
		micro.Version("v1"),
	)

	// 链路追踪
	t, closer, err := tracer.NewJaegerTracer("go.micro.api.console", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(client.AccountWrapper(service)),
		// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
		micro.WrapHandler(opentracing.NewHandlerWrapper(t)),         // server端handler接受请求
		micro.WrapSubscriber(opentracing.NewSubscriberWrapper(nil)), // server端subscriber接受消息
		micro.WrapClient(opentracing.NewClientWrapper(nil)),         // client端发起请求，包括Call()、Stream()、Publish()
		micro.WrapCall(opentracing.NewCallWrapper(t)),               // client端发起请求，仅Call()
	)

	// Register Handler
	handler.RegisterHandler(service.Server())

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
