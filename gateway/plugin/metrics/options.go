package metrics

import (
	"github.com/micro-in-cn/starter-kit/gateway/plugin/util/request"
)

type Options struct {
	skipperFunc request.SkipperFunc
}

type Option func(o *Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		skipperFunc: request.DefaultSkipperFunc,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func WithSkipperFunc(skipperFunc request.SkipperFunc) Option {
	return func(o *Options) {
		o.skipperFunc = skipperFunc
	}
}
