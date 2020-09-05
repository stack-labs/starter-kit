package auth

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/utils/request"
	"github.com/micro-in-cn/starter-kit/pkg/plugin/utils/response"
	"github.com/micro/go-micro/v3/logger"
)

// Options of auth
type Options struct {
	claims            jwt.Claims
	claimsSubjectFunc func(claims jwt.Claims) string
	headerFunc        func(r *http.Request, claims jwt.Claims)
	responseHandler   response.Handler
	skipperFunc       request.SkipperFunc
}

// Option of auth
type Option func(o *Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		claims: &jwt.StandardClaims{},
		claimsSubjectFunc: func(claims jwt.Claims) string {
			c := claims.(*jwt.StandardClaims)
			return c.Subject
		},
		headerFunc: func(r *http.Request, claims jwt.Claims) {
			c := claims.(*jwt.StandardClaims)
			// 将Claims ID加入Header
			r.Header.Set("User-Id", c.Id)
			logger.Errorf("userId: %v", c.Id)
		},
		responseHandler: response.DefaultResponseHandler,
		skipperFunc:     request.DefaultSkipperFunc,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// 自定义jwt claims
func WithClaims(claims jwt.Claims) Option {
	return func(o *Options) {
		o.claims = claims
	}
}

// 获取claims subject，Casbin权限验证的角色
func WithClaimsSubjectFunc(f func(claims jwt.Claims) string) Option {
	return func(o *Options) {
		o.claimsSubjectFunc = f
	}
}

// 添加Header
func WithHeaderFunc(f func(r *http.Request, claims jwt.Claims)) Option {
	return func(o *Options) {
		o.headerFunc = f
	}
}

// WithResponseHandler of auth
func WithResponseHandler(handler response.Handler) Option {
	return func(o *Options) {
		o.responseHandler = handler
	}
}

// WithSkipperFunc of auth
func WithSkipperFunc(skipperFunc request.SkipperFunc) Option {
	return func(o *Options) {
		o.skipperFunc = skipperFunc
	}
}
