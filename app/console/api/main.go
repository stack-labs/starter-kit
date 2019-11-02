package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"

	"github.com/micro-in-cn/starter-kit/app/console/api/client"
	"github.com/micro-in-cn/starter-kit/app/console/api/handler"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.console"),
		micro.Version("v1"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(client.AccountWrapper(service)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(nil)),       // server端handler接受请求
		micro.WrapSubscriber(opentracing.NewSubscriberWrapper(nil)), // server端subscriber接受消息
		micro.WrapClient(opentracing.NewClientWrapper(nil)),         // client端发起请求，包括Call()、Stream()、Publish()
		micro.WrapCall(opentracing.NewCallWrapper(nil)),             // client端发起请求，仅Call()
	)

	// Register Handler
	handler.RegisterHandler(service.Server())

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
