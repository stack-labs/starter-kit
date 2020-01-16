package handler

import (
	"github.com/micro/go-micro/server"
)

func RegisterHandler(server server.Server) {
	registerAccount(server)
}
