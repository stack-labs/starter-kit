package auth

import (
	"crypto/rsa"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go/test"
	"github.com/stack-labs/stack-rpc-plugins/service/gateway/plugin"
	"github.com/stack-labs/stack-rpc/config"
	"github.com/stack-labs/stack-rpc/errors"
	"github.com/stack-labs/stack-rpc/pkg/cli"
	"github.com/stack-labs/stack-rpc/util/log"
)

const id = "stack.rpc.gateway.auth"

var adapters map[string]persist.Adapter
var watchers map[string]persist.Watcher

func init() {
	adapters = make(map[string]persist.Adapter)
	watchers = make(map[string]persist.Watcher)
}

//RegisterAdapter of auth
func RegisterAdapter(key string, a persist.Adapter) {
	adapters[key] = a
}

//RegisterWatcher of auth
func RegisterWatcher(key string, w persist.Watcher) {
	watchers[key] = w
}

type authConfig struct {
	PublicKey     string `json:"public_key"`
	PublicUser    string `json:"public_user"`
	CasbinModel   string `json:"casbin_model"`
	CasbinAdapter string `json:"casbin_adapter"`
	CasbinWatcher string `json:"casbin_watcher"`
}

//Auth for micro
type Auth struct {
	options Options

	enforcer *casbin.Enforcer
	pubUser  string
	pubKey   *rsa.PublicKey
}

func (a *Auth) keyFunc(*jwt.Token) (interface{}, error) {
	return a.pubKey, nil
}

func (a *Auth) handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if a.options.skipperFunc(r) {
			h.ServeHTTP(w, r)
			return
		}

		path := r.URL.Path
		method := r.Method

		log.Infof("path: %s, method: %s, pubUser: %s", path, method, a.pubUser)

		// Public接口
		if a.pubUser != "" {
			allowed, err := a.enforcer.Enforce(a.pubUser, path, method)
			if err != nil {
				a.options.responseHandler(w, r, errors.InternalServerError(id, err.Error()))
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
			a.options.responseHandler(w, r, errors.Unauthorized(id, "JWT token parse token=nil or with error: %v", err.Error()))
			return
		}

		if !token.Valid {
			a.options.responseHandler(w, r, errors.Unauthorized(id, "JWT token invalid"))
			return
		}

		// 访问控制
		claims := token.Claims.(*jwt.StandardClaims)
		if allowed, err := a.enforcer.Enforce(claims.Id, path, method); err != nil {
			a.options.responseHandler(w, r, errors.InternalServerError(id, err.Error()))
			return
		} else if !allowed {
			log.Infof("Claims ID: %v, path: %v, method: %v", claims.Id, path, method)
			a.options.responseHandler(w, r, errors.Forbidden(id, "Casbin access control not allowed"))
			return
		}

		// 将用户ID加入Header
		r.Header.Set("UserId", claims.Id)

		// otherwise serve everything
		h.ServeHTTP(w, r)
	})
}

//NewPlugin for auth
func NewPlugin(opts ...Option) plugin.Plugin {
	options := newOptions(opts...)

	a := &Auth{
		options: options,
	}

	var egAdapter, egWatcher []string
	for k := range adapters {
		egAdapter = append(egAdapter, k)
	}
	for k := range watchers {
		egWatcher = append(egWatcher, k)
	}

	return plugin.NewPlugin(
		plugin.WithName("Auth"),
		plugin.WithFlag(
			cli.StringFlag{
				Name:  "gateway-auth-public_key",
				Usage: "Auth public key file",
				Value: "./conf/auth_key.pub",
			},
			cli.StringFlag{
				Name:  "gateway-auth-public_user",
				Usage: "Auth public user",
				Value: "public",
			},
			cli.StringFlag{
				Name:  "gateway-auth-casbin_model",
				Usage: "Casbin model config file",
				Value: "./conf/casbin_model.conf",
			},
			cli.StringFlag{
				Name:  "gateway-auth-casbin_adapter",
				Usage: "Casbin registed adapter {" + strings.Join(egAdapter, ", ") + "}",
				Value: "default",
			},
			cli.StringFlag{
				Name:  "gateway-auth-casbin_watcher",
				Usage: "Casbin registed watcher {" + strings.Join(egWatcher, ", ") + "}",
				Value: "default",
			},
		),
		plugin.WithHandler(a.handler),
		plugin.WithInit(func(cfg config.Config) error {
			conf := &authConfig{}
			cfg.Get(options.configPaths...).Scan(conf)

			// TODO config public key with config source
			pubKey := conf.PublicKey
			a.pubKey = test.LoadRSAPublicKeyFromDisk(pubKey)
			a.pubUser = conf.PublicUser

			model := conf.CasbinModel
			adapter := conf.CasbinAdapter
			watcher := conf.CasbinWatcher

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
