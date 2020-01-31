package main

import (
	"github.com/micro-in-cn/starter-kit/hipstershop/recommendationservice/handler"
	"github.com/micro-in-cn/starter-kit/hipstershop/recommendationservice/subscriber"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	recommendationservice "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.hipstershop.recommendation"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	recommendationservice.RegisterRecommendationServiceHandler(service.Server(), new(handler.RecommendationService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.recommendation", service.Server(), new(subscriber.RecommendationService))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.recommendation", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
