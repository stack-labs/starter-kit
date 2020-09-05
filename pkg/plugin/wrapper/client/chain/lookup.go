package chain

import (
	"context"
	"sort"
	"strings"

	"github.com/micro/go-micro/v3/client"
	"github.com/micro/go-micro/v3/errors"
	"github.com/micro/go-micro/v3/metadata"
	"github.com/micro/go-micro/v3/router"
)

type chainLookup struct {
	opts Options
}

// LookupRoute for a request using the router and then choose one using the selector
func (c *chainLookup) LookupRoute(ctx context.Context, req client.Request, opts client.CallOptions) ([]string, error) {
	// check to see if an address was provided as a call option
	if len(opts.Address) > 0 {
		return opts.Address, nil
	}

	// construct the router query
	query := []router.LookupOption{}

	// if a custom network was requested, pass this to the router. By default the router will use it's
	// own network, which is set during initialisation.
	if len(opts.Network) > 0 {
		query = append(query, router.LookupNetwork(opts.Network))
	}

	// lookup the routes which can be used to execute the request
	routes, err := opts.Router.Lookup(req.Service(), query...)
	if err == router.ErrRouteNotFound {
		return nil, errors.InternalServerError("go.micro.client", "service %s: %s", req.Service(), err.Error())
	} else if err != nil {
		return nil, errors.InternalServerError("go.micro.client", "error getting next %s node: %s", req.Service(), err.Error())
	}

	// 流量染色
	if val, ok := metadata.Get(ctx, c.opts.chainKey); ok && len(val) > 0 {
		chains := strings.Split(val, c.opts.chainSep)
		routes = c.filterChain(chains, routes)
	}

	// sort by lowest metric first
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Metric < routes[j].Metric
	})

	var addrs []string

	for _, route := range routes {
		addrs = append(addrs, route.Address)
	}

	return addrs, nil
}

func (c *chainLookup) filterChain(chains []string, old []router.Route) []router.Route {
	var routes []router.Route

	chain := ""
	idx := 0
	for _, route := range old {
		if route.Metadata == nil {
			continue
		}

		val := route.Metadata[c.opts.labelKey]
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

// TODO 通过自定义 Client 的 LookupFunc 来支持 Route 筛选，等官方更好的支持
// https://github.com/micro/go-micro/issues/1853
func NewClientLookup(opts ...Option) client.LookupFunc {
	options := newOptions(opts...)
	cl := &chainLookup{
		opts: options,
	}

	return cl.LookupRoute
}
