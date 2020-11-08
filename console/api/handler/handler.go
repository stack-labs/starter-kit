package handler

import (
	"github.com/micro/go-micro/v3/server"
	"github.com/micro/go-micro/v3/util/log"
)

func RegisterHandler(server server.Server) {
	if err := registerAccount(server); err != nil {
		log.Error(err)
	}
}
