package cors

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/micro/micro/plugin"
	"github.com/rs/cors"
)

// Headers
const (
	HeaderVary   = "Vary"
	HeaderOrigin = "Origin"

	// Access control
	HeaderAccessControlRequestMethod    = "Access-Control-Request-Method"
	HeaderAccessControlRequestHeaders   = "Access-Control-Request-Headers"
	HeaderAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods     = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	HeaderAccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge           = "Access-Control-Max-Age"
)

func newPluginRS(options Options) plugin.Plugin {
	return plugin.NewPlugin(
		plugin.WithHandler(func(h http.Handler) http.Handler {
			hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				h.ServeHTTP(w, r)
			})

			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				cors.New(cors.Options{
					AllowOriginFunc: func(origin string) bool {
						return true
					},
					AllowedMethods:   options.allowMethods,
					AllowedHeaders:   []string{"*"},
					AllowCredentials: options.allowCredentials,
					MaxAge:           options.maxAge,
				}).ServeHTTP(w, r, hf)
			})
		}),
	)

}

func newPlugin(options Options) plugin.Plugin {
	allowMethods := "*"
	exposeHeaders := ""
	if len(options.allowMethods) > 0 {
		allowMethods = strings.Join(options.allowMethods, ", ")
	}
	if len(options.exposeHeaders) > 0 {
		exposeHeaders = strings.Join(options.exposeHeaders, ", ")
	}

	return plugin.NewPlugin(
		plugin.WithHandler(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				origin := r.Header.Get(HeaderOrigin)
				headers := r.Header.Get(HeaderAccessControlRequestHeaders)

				if origin != "" {
					w.Header().Set(HeaderAccessControlAllowOrigin, origin)
				} else {
					w.Header().Set(HeaderAccessControlAllowOrigin, "*")
				}

				if options.allowCredentials {
					w.Header().Set(HeaderAccessControlAllowCredentials, "true")
				}

				if r.Method == http.MethodOptions && headers != "" {
					w.Header().Add(HeaderVary, HeaderOrigin)
					w.Header().Add(HeaderVary, HeaderAccessControlRequestMethod)
					w.Header().Add(HeaderVary, HeaderAccessControlRequestHeaders)
					w.Header().Set(HeaderAccessControlAllowHeaders, headers)
					w.Header().Set(HeaderAccessControlAllowMethods, allowMethods)

					if options.maxAge > 0 {
						w.Header().Set(HeaderAccessControlMaxAge, strconv.Itoa(options.maxAge))
					}

					w.WriteHeader(http.StatusNoContent)
					return
				}

				w.Header().Add(HeaderVary, HeaderOrigin)
				if exposeHeaders != "" {
					w.Header().Set(HeaderAccessControlExposeHeaders, exposeHeaders)
				}

				h.ServeHTTP(w, r)
			})
		}),
	)
}

func NewPlugin(opts ...Option) plugin.Plugin {
	options := newOptions(opts...)
	if options.useRsPkg {
		return newPluginRS(options)
	}

	return newPlugin(options)
}
