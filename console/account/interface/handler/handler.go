package handler

import (
	"github.com/micro/go-micro/v2/server"

	pb "github.com/micro-in-cn/starter-kit/console/account/genproto/srv"
	"github.com/micro-in-cn/starter-kit/console/account/registry"
	"github.com/micro-in-cn/starter-kit/console/account/usecase"
)

func Apply(server server.Server, ctn *registry.Container) {
	pb.RegisterAccountHandler(server, NewAccountService(ctn.Resolve("user-usecase").(usecase.UserUsecase)))
}
