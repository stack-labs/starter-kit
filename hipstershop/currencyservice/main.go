package main

import (
	"github.com/micro-in-cn/starter-kit/hipstershop/currencyservice/handler"
	"github.com/micro-in-cn/starter-kit/hipstershop/currencyservice/subscriber"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	currencyservice "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.hipstershop.currency"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	currencyservice.RegisterCurrencyServiceHandler(service.Server(), new(handler.CurrencyService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.currency", service.Server(), new(subscriber.CurrencyService))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.currency", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
