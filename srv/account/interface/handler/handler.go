package handler

import (
	"github.com/micro/go-micro/server"

	"github.com/micro-in-cn/starter-kit/srv/account/registry"
	"github.com/micro-in-cn/starter-kit/srv/account/usecase"
	pb "github.com/micro-in-cn/starter-kit/srv/pb/account"
)

func Apply(server server.Server, ctn *registry.Container) {
	pb.RegisterAccountHandler(server, NewAccountService(ctn.Resolve("user-usecase").(usecase.UserUsecase)))
}
