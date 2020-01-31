package chain

import (
	"net/http"
	"strings"

	"github.com/micro/cli/v2"
	"github.com/micro/micro/v2/plugin"
)

type chainPlugin struct {
	opts Options
}

func (*chainPlugin) Flags() []cli.Flag {
	return nil
}

func (*chainPlugin) Commands() []*cli.Command {
	return nil
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

func (*chainPlugin) Init(ctx *cli.Context) error {
	return nil
}

func (*chainPlugin) String() string {
	return "chain"
}

func New(opts ...Option) plugin.Plugin {
	options := newOptions(opts...)
	return &chainPlugin{opts: options}
}
