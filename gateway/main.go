package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/micro/v2/gateway/cmd"
	"github.com/micro/micro/v2/gateway/router"

	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/client/router_filter"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/select/chain"
)

func main() {
	cmd.Init(
		// 流量染色
		// TODO micro默认的api和web网关均不支持服务筛选，需要自己改造
		// https://micro.mu/blog/cn/2019/12/09/go-micro-service-chain.html
		// 这个方案不可取，查考 asim 在 pull#1388 给的反馈意见，
		// https://github.com/micro/go-micro/pull/1388
		// 自定义 Router 参考 fork 的分支版本
		// https://github.com/hb-chen/micro/tree/gateway/gateway
		// Router services filter
		router.WithOption(
			router.WithFilter(chain.NewRouterFilter()),
		),
		// 路由筛选
		micro.WrapCall(router_filter.NewCallWrapper()),
		micro.AfterStop(pluginAfterFunc),
	)
}
