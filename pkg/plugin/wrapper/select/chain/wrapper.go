package chain

import (
	"context"
	"strings"

	"github.com/micro/go-micro/v3/client"
	"github.com/micro/go-micro/v3/metadata"
	"github.com/micro/go-micro/v3/router"
	"github.com/micro/go-micro/v3/selector"
)

type chainWrapper struct {
	opts Options
	client.Client
}

func (w *chainWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	if val, ok := metadata.Get(ctx, w.opts.chainKey); ok && len(val) > 0 {
		chains := strings.Split(val, w.opts.chainSep)
		nOpts := append(opts, client.WithSelectOptions(
			selector.WithFilter(w.filterChain(chains)),
		))
		return w.Client.Call(ctx, req, rsp, nOpts...)
	}

	return w.Client.Call(ctx, req, rsp, opts...)
}

func (w *chainWrapper) filterChain(chains []string) selector.Filter {
	return func(old []router.Route) []router.Route {
		var routes []router.Route

		chain := ""
		idx := 0
		for _, route := range old {
			if route.Metadata == nil {
				continue
			}

			val := route.Metadata[w.opts.labelKey]
			if len(val) == 0 {
				continue
			}

			if len(chain) > 0 && idx == 0 {
				if chain == val {
					routes = append(routes, route)
				}
				continue
			}

			// chains按顺序优先匹配
			ok, i := inArray(val, chains)
			if ok && idx > i {
				// 出现优先链路，services清空，nodes清空
				idx = i
				routes = routes[:0]
			}

			if ok {
				chain = val
				routes = append(routes, route)
			}
		}

		if len(routes) == 0 {
			return old
		}

		return routes
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
