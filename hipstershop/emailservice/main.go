package main

import (
	"github.com/micro-in-cn/starter-kit/hipstershop/emailservice/handler"
	"github.com/micro-in-cn/starter-kit/hipstershop/emailservice/subscriber"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	emailservice "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.hipstershop.email"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	emailservice.RegisterEmailServiceHandler(service.Server(), new(handler.EmailService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.email", service.Server(), new(subscriber.EmailService))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.email", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
