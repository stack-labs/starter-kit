package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	"github.com/micro-in-cn/starter-kit/srv/account/interface/handler"
	"github.com/micro-in-cn/starter-kit/srv/account/registry"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.account"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	ctn, err := registry.NewContainer()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	// Register Handler
	handler.Apply(service.Server(), ctn)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
