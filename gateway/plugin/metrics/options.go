package metrics

import (
	"github.com/micro-in-cn/starter-kit/gateway/plugin/util/request"
)

type Options struct {
	namespace string
	subsystem string

	skipperFunc request.SkipperFunc
}

type Option func(o *Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		namespace:   "http",
		subsystem:   "",
		skipperFunc: request.DefaultSkipperFunc,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func WithNamespace(namespace string) Option {
	return func(o *Options) {
		o.namespace = namespace
	}
}

func WithSubsystem(subsystem string) Option {
	return func(o *Options) {
		o.subsystem = subsystem
	}
}

func WithSkipperFunc(skipperFunc request.SkipperFunc) Option {
	return func(o *Options) {
		o.skipperFunc = skipperFunc
	}
}
