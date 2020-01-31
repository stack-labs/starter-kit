package main

import (
	"github.com/micro-in-cn/starter-kit/hipstershop/productcatalogservice/handler"
	"github.com/micro-in-cn/starter-kit/hipstershop/productcatalogservice/subscriber"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	productcatalogservice "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.hipstershop.productcatalog"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	productcatalogservice.RegisterProductCatalogServiceHandler(service.Server(), new(handler.ProductcatalogService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.productcatalog", service.Server(), new(subscriber.ProductcatalogService))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.productcatalog", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
