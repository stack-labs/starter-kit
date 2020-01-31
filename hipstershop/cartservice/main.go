package main

import (
	"github.com/micro-in-cn/starter-kit/hipstershop/cartservice/handler"
	"github.com/micro-in-cn/starter-kit/hipstershop/cartservice/subscriber"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	cartservice "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.hipstershop.cart"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	cartservice.RegisterCartServiceHandler(service.Server(), new(handler.CartService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.cart", service.Server(), new(subscriber.CartService))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.cart", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
