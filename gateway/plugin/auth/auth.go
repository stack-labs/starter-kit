package auth

import (
	"crypto/rsa"
	"github.com/micro/go-micro/util/log"

	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go/test"
	"github.com/micro/cli"
	"github.com/micro/go-micro/errors"
	"github.com/micro/micro/plugin"
)

const id = "hb-go.micro-plugins.micro.auth"

var AuthResponse ResponseHandler
var adapters map[string]persist.Adapter
var watchers map[string]persist.Watcher

func init() {
	AuthResponse = DefaultResponseHandler
	adapters = make(map[string]persist.Adapter)
	watchers = make(map[string]persist.Watcher)
}

func RegisterAdapter(key string, a persist.Adapter) {
	adapters[key] = a
}

func RegisterWatcher(key string, w persist.Watcher) {
	watchers[key] = w
}

type ResponseHandler func(w http.ResponseWriter, r *http.Request, err error)

func DefaultResponseHandler(w http.ResponseWriter, r *http.Request, err error) {
	mErr, ok := err.(*errors.Error)
	if !ok {
		mErr = &errors.Error{
			Id:     id,
			Code:   http.StatusInternalServerError,
			Detail: err.Error(),
			Status: http.StatusText(http.StatusInternalServerError),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(mErr.Code))
	w.Write([]byte(mErr.Error()))
}

type Auth struct {
	opts Options

	enforcer *casbin.Enforcer
	pubUser  string
	pubKey   *rsa.PublicKey
	response ResponseHandler
}

func (a *Auth) keyFunc(*jwt.Token) (interface{}, error) {
	return a.pubKey, nil
}

func (a *Auth) handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if a.opts.skipperFunc(r) {
			h.ServeHTTP(w, r)
			return
		}

		path := r.URL.Path
		method := r.Method

		// Public接口
		if a.pubUser != "" {
			allowed, err := a.enforcer.Enforce(a.pubUser, path, method)
			if err != nil {
				a.response(w, r, errors.InternalServerError(id, err.Error()))
				return
			} else if allowed {
				h.ServeHTTP(w, r)
				return
			}
		}

		// JWT token验证
		token, err := request.ParseFromRequest(
			r,
			request.AuthorizationHeaderExtractor,
			a.keyFunc,
			request.WithClaims(&jwt.StandardClaims{}),
		)

		if err != nil || token == nil {
			a.response(w, r, errors.Unauthorized(id, "JWT token parse token=nil or with error: %v", err.Error()))
			return
		}

		if !token.Valid {
			a.response(w, r, errors.Unauthorized(id, "JWT token invalid"))
			return
		}

		// 访问控制
		claims := token.Claims.(*jwt.StandardClaims)
		if allowed, err := a.enforcer.Enforce(claims.Id, path, method); err != nil {
			a.response(w, r, errors.InternalServerError(id, err.Error()))
			return
		} else if !allowed {
			log.Infof("Claims ID: %v, path: %v, method: %v", claims.Id, path, method)
			a.response(w, r, errors.Forbidden(id, "Casbin access control not allowed"))
			return
		}

		// otherwise serve everything
		h.ServeHTTP(w, r)
	})
}

func NewPlugin(opts ...Option) plugin.Plugin {
	options := newOptions(opts...)

	a := Auth{
		opts:     options,
		response: AuthResponse,
	}

	var egAdapter, egWatcher []string
	for k, _ := range adapters {
		egAdapter = append(egAdapter, k)
	}
	for k, _ := range watchers {
		egWatcher = append(egWatcher, k)
	}

	return plugin.NewPlugin(
		plugin.WithName("Auth"),
		plugin.WithFlag(
			cli.StringFlag{
				Name:  "auth_pub_key",
				Usage: "Auth public key file",
				Value: "./conf/auth_key.pub",
			},
			cli.StringFlag{
				Name:  "casbin_model",
				Usage: "Casbin model config file",
				Value: "./conf/casbin_model.conf",
			},
			cli.StringFlag{
				Name:  "casbin_adapter",
				Usage: "Casbin registed adapter {" + strings.Join(egAdapter, ", ") + "}",
				Value: "default",
			},
			cli.StringFlag{
				Name:  "casbin_watcher",
				Usage: "Casbin registed watcher {" + strings.Join(egWatcher, ", ") + "}",
				Value: "default",
			},
			cli.StringFlag{
				Name:  "casbin_public_user",
				Usage: "Casbin public user",
				Value: "public",
			},
		),
		plugin.WithHandler(a.handler),
		plugin.WithInit(func(ctx *cli.Context) error {
			a.pubUser = ctx.String("casbin_public_user")
			pubKey := ctx.String("auth_pub_key")
			a.pubKey = test.LoadRSAPublicKeyFromDisk(pubKey)

			model := ctx.String("casbin_model")
			adapter := ctx.String("casbin_adapter")
			watcher := ctx.String("casbin_watcher")

			var e *casbin.Enforcer
			if a, ok := adapters[adapter]; ok {
				var err error
				e, err = casbin.NewEnforcer(model, a)
				if err != nil {
					return err
				}
			} else {
				return errors.New(id, "adapter not exist", http.StatusInternalServerError)
			}

			// Load the policy.
			e.LoadPolicy()

			// Set the watcher for the enforcer.
			if w, ok := watchers[watcher]; ok {
				e.SetWatcher(w)

				// Set callback to local example
				// c.watcher.SetUpdateCallback(func(string) { e.LoadPolicy() })
			}

			a.enforcer = e

			return nil
		}),
	)
}
