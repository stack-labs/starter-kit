package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/micro/cmd"
)

func main() {
	cmd.Init(
		micro.AfterStop(pluginAfterFunc),
	)
}
