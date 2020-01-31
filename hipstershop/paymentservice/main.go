package main

import (
	"github.com/micro-in-cn/starter-kit/hipstershop/paymentservice/handler"
	"github.com/micro-in-cn/starter-kit/hipstershop/paymentservice/subscriber"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	paymentservice "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.hipstershop.payment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	paymentservice.RegisterPaymentServiceHandler(service.Server(), new(handler.PaymentService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.payment", service.Server(), new(subscriber.PaymentService))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.payment", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
