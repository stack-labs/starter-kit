package client

import (
	"context"

	"github.com/stack-labs/stack-rpc/service"
	account "github.com/stack-labs/starter-kit/console/api/genproto/srv"

	"github.com/stack-labs/stack-rpc/server"
)

type exampleKey struct{}

// FromContext retrieves the client from the Context
func AccountFromContext(ctx context.Context) (account.AccountService, bool) {
	c, ok := ctx.Value(exampleKey{}).(account.AccountService)
	return c, ok
}

// Client returns a wrapper for the ExampleClient
func AccountWrapper(service service.Service) server.HandlerWrapper {
	client := account.NewAccountService("stack.rpc.srv.account", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, exampleKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
