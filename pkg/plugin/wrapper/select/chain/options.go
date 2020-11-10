package chain

type Options struct {
	chainKey   string
	chainSep   string
	labelKey   string
	chainLabel string
}

type Option func(o *Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		chainKey: "X-Micro-Chain",
		chainSep: ";",
		labelKey: "chain",
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

func WithChainLabel(label string) Option {
	return func(o *Options) {
		o.chainLabel = label
	}
}
