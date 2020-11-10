package chain

import (
	"net/http"
)

type Options struct {
	chainKey   string
	chainSep   string
	chainsFunc ChainsFunc
}

type ChainsFunc func(r *http.Request) []string
type Option func(o *Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		chainKey: "X-Micro-Chain",
		chainSep: ";",
		chainsFunc: func(r *http.Request) []string {
			return nil
		},
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func WithChainKey(k string) Option {
	return func(o *Options) {
		o.chainKey = k
	}
}

func WithChainSep(sep string) Option {
	return func(o *Options) {
		o.chainSep = sep
	}
}

func WithChainsFunc(f ChainsFunc) Option {
	return func(o *Options) {
		o.chainsFunc = f
	}
}
