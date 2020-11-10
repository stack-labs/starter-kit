package handler

import (
	"net/http"

	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/api/handler"
	"github.com/stack-labs/stack-rpc/api/handler/event"
	"github.com/stack-labs/stack-rpc/api/router"
	"github.com/stack-labs/stack-rpc/errors"

	// TODO: only import handler package
	aapi "github.com/stack-labs/stack-rpc/api/handler/api"
	ahttp "github.com/stack-labs/stack-rpc/api/handler/http"
	arpc "github.com/stack-labs/stack-rpc/api/handler/rpc"
	aweb "github.com/stack-labs/stack-rpc/api/handler/web"
)

type metaHandler struct {
	s stack.Service
	r router.Router
}

func (m *metaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	service, err := m.r.Route(r)
	if err != nil {
		er := errors.InternalServerError(m.r.Options().Namespace, err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte(er.Error()))
		return
	}

	// TODO: don't do this ffs
	switch service.Endpoint.Handler {
	// web socket handler
	case aweb.Handler:
		aweb.WithService(service, handler.WithService(m.s)).ServeHTTP(w, r)
	// proxy handler
	case "proxy", ahttp.Handler:
		ahttp.WithService(service, handler.WithService(m.s)).ServeHTTP(w, r)
	// rpcx handler
	case arpc.Handler:
		arpc.WithService(service, handler.WithService(m.s)).ServeHTTP(w, r)
	// event handler
	case event.Handler:
		ev := event.NewHandler(
			handler.WithNamespace(m.r.Options().Namespace),
			handler.WithService(m.s),
		)
		ev.ServeHTTP(w, r)
	// api handler
	case aapi.Handler:
		aapi.WithService(service, handler.WithService(m.s)).ServeHTTP(w, r)
	// default handler: rpc
	default:
		arpc.WithService(service, handler.WithService(m.s)).ServeHTTP(w, r)
	}
}

// Meta is a http.Handler that routes based on endpoint metadata
func Meta(s stack.Service, r router.Router) http.Handler {
	return &metaHandler{
		s: s,
		r: r,
	}
}
