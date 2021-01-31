package chain

import (
	"net/http"
	"strings"

	"github.com/stack-labs/stack-rpc-plugins/service/stackway/plugin"
)

type chainPlugin struct {
	opts Options
}

func (p *chainPlugin) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			chains := p.opts.chainsFunc(r)
			if len(chains) > 0 {
				join := strings.Join(chains, p.opts.chainSep)
				r.Header.Set(p.opts.chainKey, join)
			}

			h.ServeHTTP(w, r)
		})
	}
}

func New(opts ...Option) plugin.Plugin {
	options := newOptions(opts...)
	cp := &chainPlugin{opts: options}
	return plugin.NewPlugin(
		plugin.WithName("chain"),
		plugin.WithHandler(cp.Handler()),
	)
}
