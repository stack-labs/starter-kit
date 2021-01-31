package main

import (
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc-plugins/service/stackway/api"
	"github.com/stack-labs/stack-rpc/util/log"
)

func main() {
	svc := stack.NewService()

	// stackway server
	apiServer := api.NewServer(svc)
	svc.Init(apiServer.Options()...)

	// run service
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
