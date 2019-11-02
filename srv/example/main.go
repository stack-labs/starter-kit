package main

import (
	"github.com/micro-in-cn/starter-kit/srv/example/handler"
	example "github.com/micro-in-cn/starter-kit/srv/example/proto/example"
	"github.com/micro-in-cn/starter-kit/srv/example/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.example"),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry()),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.example", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.example", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
