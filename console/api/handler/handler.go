package handler

import (
	"github.com/micro/go-micro/v2/server"
)

func RegisterHandler(server server.Server) {
	registerAccount(server)
}
