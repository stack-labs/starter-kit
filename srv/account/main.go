package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"

	tracer "github.com/micro-in-cn/starter-kit/pkg/opentracing"
	"github.com/micro-in-cn/starter-kit/srv/account/interface/handler"
	"github.com/micro-in-cn/starter-kit/srv/account/registry"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.account"),
		micro.Version("latest"),
	)

	// 链路追踪
	t, closer, err := tracer.NewJaegerTracer("go.micro.srv.account", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()
	service.Init(
		micro.WrapCall(opentracing.NewCallWrapper(t)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(t)),
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
