package main

import (
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc-plugins/service/gateway"
	"github.com/stack-labs/stack-rpc/plugin"
	"github.com/stack-labs/stack-rpc/server/mock"
	"github.com/stack-labs/stack-rpc/util/log"
)

func init() {
	plugin.DefaultServers["mock"] = mock.NewServer
}

func main() {
	svc := stack.NewService()

	// run gateway
	gateway.Run(svc)

	// run service
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
