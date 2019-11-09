package cors

import (
	"net/http"
	"strings"
)

type Options struct {
	allowMethods     []string
	exposeHeaders    []string
	allowCredentials bool
	maxAge           int

	useRsPkg bool
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		allowMethods:     []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		allowCredentials: true,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func WithAllowMethods(methods ...string) Option {
	return func(o *Options) {
		o.allowMethods = convert(methods, strings.ToUpper)
	}
}

func WithExposeHeaders(headers ...string) Option {
	return func(o *Options) {
		o.exposeHeaders = convert(headers, http.CanonicalHeaderKey)
	}
}

func WithAllowCredentials(allow bool) Option {
	return func(o *Options) {
		o.allowCredentials = allow
	}
}

func WithMaxAge(maxAge int) Option {
	return func(o *Options) {
		o.maxAge = maxAge
	}
}

func WithUseRsPkg(useRs bool) Option {
	return func(o *Options) {
		o.useRsPkg = useRs
	}
}

type converter func(string) string

// convert converts a list of string using the passed converter function
func convert(s []string, c converter) []string {
	out := []string{}
	for _, i := range s {
		out = append(out, c(i))
	}
	return out
}
