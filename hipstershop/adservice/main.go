package main

import (
	"github.com/micro-in-cn/starter-kit/hipstershop/adservice/handler"
	"github.com/micro-in-cn/starter-kit/hipstershop/adservice/subscriber"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	adservice "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.hipstershop.ad"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	adservice.RegisterAdServiceHandler(service.Server(), new(handler.AdService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.ad", service.Server(), new(subscriber.AdService))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.ad", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
