package chain

import (
	"context"
	"strings"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/registry"
)

type chainWrapper struct {
	opts Options
	client.Client
}

func (w *chainWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	if val, ok := metadata.Get(ctx, w.opts.chainKey); ok && len(val) > 0 {
		chains := strings.Split(val, w.opts.chainSep)
		nOpts := append(opts, client.WithSelectOption(
			selector.WithFilter(w.filterChain(chains)),
		))
		return w.Client.Call(ctx, req, rsp, nOpts...)
	}

	return w.Client.Call(ctx, req, rsp, opts...)
}

func (w *chainWrapper) filterChain(chains []string) selector.Filter {
	return func(old []*registry.Service) []*registry.Service {
		var services []*registry.Service

		chain := ""
		idx := 0
		for _, service := range old {
			serv := new(registry.Service)
			var nodes []*registry.Node

			for _, node := range service.Nodes {
				if node.Metadata == nil {
					continue
				}

				val := node.Metadata[w.opts.labelKey]
				if len(val) == 0 {
					continue
				}

				if len(chain) > 0 && idx == 0 {
					if chain == val {
						nodes = append(nodes, node)
					}
					continue
				}

				// chains按顺序优先匹配
				ok, i := inArray(val, chains)
				if ok && idx > i {
					// 出现优先链路，services清空，nodes清空
					idx = i
					services = services[:0]
					nodes = nodes[:0]
				}

				if ok {
					chain = val
					nodes = append(nodes, node)
				}
			}

			// only add service if there's some nodes
			if len(nodes) > 0 {
				// copy
				*serv = *service
				serv.Nodes = nodes
				services = append(services, serv)
			}
		}

		if len(services) == 0 {
			return old
		}

		return services
	}
}

func inArray(s string, d []string) (bool, int) {
	for k, v := range d {
		if s == v {
			return true, k
		}
	}
	return false, 0
}

func NewClientWrapper(opts ...Option) client.Wrapper {
	options := newOptions(opts...)
	return func(c client.Client) client.Client {
		return &chainWrapper{
			opts:   options,
			Client: c,
		}
	}
}
