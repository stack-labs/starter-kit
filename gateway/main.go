package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/micro/v2/cmd"

	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/select/chain"
)

func main() {
	cmd.Init(
		// 流量染色
		// TODO micro默认的api和web网关均不支持服务筛选，需要自己改造
		// https://micro.mu/blog/cn/2019/12/09/go-micro-service-chain.html
		micro.WrapClient(chain.NewClientWrapper()),
		micro.AfterStop(pluginAfterFunc),
	)
}
