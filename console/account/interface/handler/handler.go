package handler

import (
	"github.com/micro/go-micro/v3/server"
	"go.uber.org/dig"

	pb "github.com/micro-in-cn/starter-kit/console/account/genproto/srv"
	"github.com/micro-in-cn/starter-kit/console/account/usecase"
)

func Apply(server server.Server, c *dig.Container) {
	c.Invoke(func(userUsecase usecase.UserUsecase) {
		pb.RegisterAccountHandler(server, NewAccountService(userUsecase))
	})
}
