package handler

import (
	"github.com/stack-labs/stack-rpc/server"

	pb "github.com/stack-labs/starter-kit/console/account/genproto/srv"
	"github.com/stack-labs/starter-kit/console/account/registry"
	"github.com/stack-labs/starter-kit/console/account/usecase"
)

func Apply(server server.Server, ctn *registry.Container) {
	pb.RegisterAccountHandler(server, NewAccountService(ctn.Resolve("user-usecase").(usecase.UserUsecase)))
}
