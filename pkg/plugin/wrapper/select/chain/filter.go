package chain

import (
	"net/http"
	"strings"

	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/micro/v2/gateway/router"
)

type chainFilter struct {
	opts Options
}

func (f *chainFilter) routerFilter(req *http.Request) selector.Filter {
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
