package validate

import (
	"context"
	"fmt"
	"reflect"

	"github.com/micro/go-micro/v3/client"
	"github.com/micro/go-micro/v3/errors"
	"github.com/micro/go-micro/v3/server"
)

func NewCallWrapper(opts ...Option) client.CallWrapper {
	o := newOptions(opts)
	return func(callFunc client.CallFunc) client.CallFunc {
		return func(ctx context.Context, addr string, req client.Request, rsp interface{}, opts client.CallOptions) error {
			t := reflect.TypeOf(req.Body())
			if m, ok := t.MethodByName(o.funcName); ok {
				if e := m.Func.Call([]reflect.Value{reflect.ValueOf(req.Body())}); len(e) > 0 {
					err := e[0].Interface()
					if err != nil {
						return errors.BadRequest(
							"call.wrapper.validator",
							fmt.Sprintf("service: %s method: %s error: %v", req.Service(), req.Method(), err))
					}
				}
			}

			return callFunc(ctx, addr, req, rsp, opts)
		}
	}
}

func NewHandlerWrapper(opts ...Option) server.HandlerWrapper {
	o := newOptions(opts)
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			t := reflect.TypeOf(req.Body())
			if m, ok := t.MethodByName(o.funcName); ok {
				if e := m.Func.Call([]reflect.Value{reflect.ValueOf(req.Body())}); len(e) > 0 {
					err := e[0].Interface()
					if err != nil {
						return errors.BadRequest(
							"handler.wrapper.validator",
							fmt.Sprintf("service: %s method: %s error: %v", req.Service(), req.Method(), err))
					}
				}
			}

			return handlerFunc(ctx, req, rsp)
		}
	}
}
