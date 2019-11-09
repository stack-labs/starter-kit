package auth

import (
	"net/http"
)

type Options struct {
	skipperFunc SkipperFunc
}

type Option func(o *Options)
type SkipperFunc func(r *http.Request) bool

var DefaultSkipperFunc = func(r *http.Request) bool {
	return false
}

func newOptions(opts ...Option) Options {
	opt := Options{
		skipperFunc: DefaultSkipperFunc,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func WithSkipperFunc(skipperFunc SkipperFunc) Option {
	return func(o *Options) {
		o.skipperFunc = skipperFunc
	}
}
