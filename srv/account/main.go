package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	
	tracer "github.com/micro-in-cn/starter-kit/pkg/opentracing"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/select/chain"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/trace/opentracing"
	"github.com/micro-in-cn/starter-kit/srv/account/interface/handler"
	"github.com/micro-in-cn/starter-kit/srv/account/registry"
)

func init() {
	// 配置加载
	err := config.LoadFile("./conf/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	md := make(map[string]string)
	md["chain"] = "gray"

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.account"),
		micro.Version("latest"),
		micro.Metadata(md),
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
