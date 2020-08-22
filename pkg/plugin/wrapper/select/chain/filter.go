package chain

import (
	"net/http"
	"strings"

	"github.com/micro/go-micro/v3/registry"
	"github.com/micro/micro/v3/client/gateway/router"
)

type chainFilter struct {
	opts Options
}

func (f *chainFilter) routerFilter(req *http.Request) router.ServiceFilter {
	if val := req.Header.Get(f.opts.chainKey); len(val) > 0 {
		chains := strings.Split(val, f.opts.chainSep)
		return filterChain(f.opts.labelKey, chains)
	}

	return func(services []*registry.Service) []*registry.Service {
		return services
	}
}

func NewRouterFilter(opts ...Option) router.Filter {
	options := newOptions(opts...)
	w := &chainFilter{
		opts: options,
	}

	return w.routerFilter
}
