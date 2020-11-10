package handler

import (
	"github.com/stack-labs/stack-rpc/server"
)

func RegisterHandler(server server.Server) {
	registerAccount(server)
}
