package main

import (
	"github.com/micro-in-cn/starter-kit/hipstershop/shippingservice/handler"
	"github.com/micro-in-cn/starter-kit/hipstershop/shippingservice/subscriber"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	shippingservice "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.hipstershop.shipping"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	shippingservice.RegisterShippingServiceHandler(service.Server(), new(handler.ShippingService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.shipping", service.Server(), new(subscriber.ShippingService))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.shipping", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
