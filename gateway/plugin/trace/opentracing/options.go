package opentracing

import (
	"github.com/opentracing/opentracing-go"

	"github.com/micro-in-cn/starter-kit/gateway/plugin/util/request"
	"github.com/micro-in-cn/starter-kit/gateway/plugin/util/response"
)

type Options struct {
	tracer opentracing.Tracer

	responseHandler response.Handler
	skipperFunc     request.SkipperFunc

	autoStart bool
}

type Option func(o *Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		tracer:          opentracing.GlobalTracer(),
		responseHandler: response.DefaultResponseHandler,
		skipperFunc:     request.DefaultSkipperFunc,
		autoStart:       true,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func WithTracer(tracer opentracing.Tracer) Option {
	return func(o *Options) {
		o.tracer = tracer
	}
}

func WithResponseHandler(handler response.Handler) Option {
	return func(o *Options) {
		o.responseHandler = handler
	}
}

func WithSkipperFunc(skipperFunc request.SkipperFunc) Option {
	return func(o *Options) {
		o.skipperFunc = skipperFunc
	}
}

func WithAutoStart(auto bool) Option {
	return func(o *Options) {
		o.autoStart = auto
	}
}
