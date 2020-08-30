package opentracing

import (
	"github.com/opentracing/opentracing-go"

	"github.com/micro-in-cn/x-gateway/utils/request"
	"github.com/micro-in-cn/x-gateway/utils/response"
)

//Options of opentracing
type Options struct {
	tracer opentracing.Tracer

	responseHandler response.Handler
	skipperFunc     request.SkipperFunc

	autoStart bool
}

//Option of tace
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

//WithTracer of opentracing
func WithTracer(tracer opentracing.Tracer) Option {
	return func(o *Options) {
		o.tracer = tracer
	}
}

//WithResponseHandler of opentracing
func WithResponseHandler(handler response.Handler) Option {
	return func(o *Options) {
		o.responseHandler = handler
	}
}

//WithSkipperFunc of opentracing
func WithSkipperFunc(skipperFunc request.SkipperFunc) Option {
	return func(o *Options) {
		o.skipperFunc = skipperFunc
	}
}

//WithAutoStart of opentracing
func WithAutoStart(auto bool) Option {
	return func(o *Options) {
		o.autoStart = auto
	}
}
