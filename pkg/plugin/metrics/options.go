package metrics

import (
	"github.com/micro-in-cn/x-gateway/utils/request"
)

//Options of metric
type Options struct {
	namespace string
	subsystem string

	skipperFunc request.SkipperFunc
}

//Option of metric
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

//WithNamespace of metric
func WithNamespace(namespace string) Option {
	return func(o *Options) {
		o.namespace = namespace
	}
}

//WithSubsystem of metric
func WithSubsystem(subsystem string) Option {
	return func(o *Options) {
		o.subsystem = subsystem
	}
}

//WithSkipperFunc of metric
func WithSkipperFunc(skipperFunc request.SkipperFunc) Option {
	return func(o *Options) {
		o.skipperFunc = skipperFunc
	}
}
