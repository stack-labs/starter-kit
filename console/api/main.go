package main

import (
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/util/log"

	"github.com/stack-labs/starter-kit/console/api/client"
	"github.com/stack-labs/starter-kit/console/api/handler"
	tracer "github.com/stack-labs/starter-kit/pkg/opentracing"
	"github.com/stack-labs/starter-kit/pkg/plugin/wrapper/select/chain"
	"github.com/stack-labs/starter-kit/pkg/plugin/wrapper/trace/opentracing"
)

func main() {
	md := make(map[string]string)
	md["chain"] = "gray"

	// New Service
	service := stack.NewService(
		stack.Name("go.micro.api.console"),
		stack.Version("v1"),
		stack.Metadata(md),
	)

	// 链路追踪
	t, closer, err := tracer.NewJaegerTracer("go.micro.api.console", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()

	// Initialise service
	service.Init(
		stack.WrapClient(chain.NewClientWrapper()),
	)
	service.Init(
		// create wrap for the Example srv client
		stack.WrapHandler(client.AccountWrapper(service)),
		// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
		stack.WrapHandler(opentracing.NewHandlerWrapper(t)),         // server端handler接受请求
		stack.WrapSubscriber(opentracing.NewSubscriberWrapper(nil)), // server端subscriber接受消息
		stack.WrapClient(opentracing.NewClientWrapper(nil)),         // client端发起请求，包括Call()、Stream()、Publish()
		stack.WrapCall(opentracing.NewCallWrapper(t)),               // client端发起请求，仅Call()
	)

	// Register Handler
	handler.RegisterHandler(service.Server())

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
