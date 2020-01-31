package main

import (
	"github.com/micro-in-cn/starter-kit/hipstershop/checkoutservice/handler"
	"github.com/micro-in-cn/starter-kit/hipstershop/checkoutservice/subscriber"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	checkoutservice "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.hipstershop.checkout"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	checkoutservice.RegisterCheckoutServiceHandler(service.Server(), new(handler.CheckoutService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.checkout", service.Server(), new(subscriber.CheckoutService))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.checkout", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
