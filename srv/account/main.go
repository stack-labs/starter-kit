package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	"github.com/micro-in-cn/starter-kit/srv/account/handler"
	account "github.com/micro-in-cn/starter-kit/srv/account/proto/account"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.account"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	account.RegisterAccountHandler(service.Server(), new(handler.Account))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
