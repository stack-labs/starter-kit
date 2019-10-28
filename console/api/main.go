package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	"github.com/micro-in-cn/starter-kit/console/api/client"
	"github.com/micro-in-cn/starter-kit/console/api/handler"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.console"),
		micro.Version("v1"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(client.AccountWrapper(service)),
	)

	// Register Handler
	handler.RegisterHandler(service.Server())

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
