package registry

import (
	"net/http"

	"github.com/micro/go-micro/v2/api/router"
	"github.com/micro/go-micro/v2/client/selector"
)

type Options struct {
	router.Options
	routerOpt []router.Option

	Filters []Filter
}

type Option func(o *Options)

type Filter func(req *http.Request) selector.Filter

func NewOptions(opts ...Option) Options {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	options.Options = router.NewOptions(options.routerOpt...)

	return options
}

func WithRouterOption(opt ...router.Option) Option {
	return func(o *Options) {
		o.routerOpt = append(o.routerOpt, opt...)
	}
}

func WithFilter(fn ...Filter) Option {
	return func(o *Options) {
		o.Filters = append(o.Filters, fn...)
	}
}
