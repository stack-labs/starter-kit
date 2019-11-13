package opentracing

import (
	"net/http"

	"github.com/micro/micro/plugin"
	"github.com/opentracing/opentracing-go"

	"github.com/micro-in-cn/starter-kit/gateway/plugin/util/response"
)

// StartSpanFromHeader returns a new span with the given operation name and options. If a span
// is found in the header, it will be used as the parent of the resulting span.
func StartSpanFromHeader(header http.Header, tracer opentracing.Tracer, name string, opts ...opentracing.StartSpanOption) (opentracing.Span, error) {

	// Find parent span.
	if spanCtx, err := tracer.Extract(opentracing.TextMap, opentracing.HTTPHeadersCarrier(header)); err == nil {
		opts = append(opts, opentracing.ChildOf(spanCtx))
	}

	sp := tracer.StartSpan(name, opts...)
	if err := sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.HTTPHeadersCarrier(header)); err != nil {
		return nil, err
	}

	return sp, nil
}

func newPlugin(opts ...Option) plugin.Plugin {
	options := newOptions(opts...)
	return plugin.NewPlugin(
		plugin.WithName("trace"),
		plugin.WithHandler(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if options.skipperFunc(r) {
					h.ServeHTTP(w, r)
					return
				}

				name := r.URL.Path
				span, err := StartSpanFromHeader(r.Header, options.tracer, name)
				if err != nil {
					options.responseHandler(w, r, err)
				}
				defer span.Finish()

				span.SetTag("http.host", r.Host)
				span.SetTag("http.method", r.Method)

				ww := response.WrapWriter{ResponseWriter: w}
				h.ServeHTTP(&ww, r)

				span.SetTag("http.status_code", ww.StatusCode)
			})
		}),
	)
}

func NewPlugin(opts ...Option) plugin.Plugin {
	return newPlugin(opts...)
}
