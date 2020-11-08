package main

import (
	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/client/chain"
	"github.com/micro/go-micro/v3/client"
	"github.com/micro/go-micro/v3/server/mock"
	"github.com/micro/go-micro/v3/util/log"
	"github.com/micro/micro/v3/client/cli/util"
	"github.com/micro/micro/v3/cmd"
	microClient "github.com/micro/micro/v3/service/client"
	"github.com/micro/micro/v3/service/gateway"
	"github.com/micro/micro/v3/service/gateway/router"
	microServer "github.com/micro/micro/v3/service/server"
	"github.com/urfave/cli/v2"

	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/client/router_filter"
	_ "github.com/micro-in-cn/starter-kit/profile"
)

const (
	// EnvLocal is a builtin environment, it represents your local `micro server`
	EnvDev = "dev"
)

var envs = map[string]util.Env{
	EnvDev: {
		Name: EnvDev,
	},
}

func init() {
	// 默认有个 EnvLocal，并且 ProxyAddress = 127.0.0.1:8081
	for _, env := range envs {
		util.AddEnv(env)
	}
}

func main() {
	// 流量染色
	// TODO micro默认的api和web网关均不支持服务筛选，需要自己改造
	// https://micro.mu/blog/cn/2019/12/09/go-micro-service-chain.html
	// 这个方案不可取，查考 asim 在 pull#1388 给的反馈意见，
	// https://github.com/micro/go-micro/pull/1388
	// 自定义 Router 参考 fork 的分支版本
	// https://github.com/hb-chen/micro/tree/gateway/service/gateway
	// Router services filter
	command := gateway.Commands(
		// 流量染色
		router.WithFilter(chain.NewRouterFilter()),
	)
	command.After = func(ctx *cli.Context) error {
		pluginAfterFunc()
		return nil
	}
	cmd.Register(command)

	microServer.DefaultServer = mock.NewServer()
	microClient.DefaultClient.Init(
		client.WrapCall(router_filter.NewCallWrapper()),
	)

	if err := cmd.DefaultCmd.Run(); err != nil {
		log.Fatal(err)
	}
}
