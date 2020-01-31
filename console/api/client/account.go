package client

import (
	"context"

	account "github.com/micro-in-cn/starter-kit/console/api/genproto/srv"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

type exampleKey struct{}

// FromContext retrieves the client from the Context
func AccountFromContext(ctx context.Context) (account.AccountService, bool) {
	c, ok := ctx.Value(exampleKey{}).(account.AccountService)
	return c, ok
}

// Client returns a wrapper for the ExampleClient
func AccountWrapper(service micro.Service) server.HandlerWrapper {
	client := account.NewAccountService("go.micro.srv.account", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, exampleKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
