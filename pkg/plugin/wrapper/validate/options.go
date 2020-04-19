package validate

type options struct {
	funcName string // 默认Validate() error，无参数，返回类型error
}

type Option func(option *options)

func newOptions(opts []Option) *options {
	opt := &options{
		funcName: "Validate",
	}
	for _, o := range opts {
		o(opt)
	}

	return opt
}

func WithFuncName(name string) Option {
	return func(option *options) {
		option.funcName = name
	}
}
