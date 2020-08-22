package router_filter

import (
	"context"
	"strings"

	"github.com/micro/go-micro/v3/client"
	"github.com/micro/go-micro/v3/metadata"
	"github.com/micro/go-micro/v3/util/log"
	microClient "github.com/micro/micro/v3/service/client"
)

func NewCallWrapper() client.CallWrapper {
	return func(callFunc client.CallFunc) client.CallFunc {
		return func(ctx context.Context, addr string, req client.Request, rsp interface{}, opts client.CallOptions) error {
			md, ok := metadata.FromContext(ctx)
			if !ok {
				md = make(map[string]string)
			}

			log.Infof("client options.proxy: %v", microClient.DefaultClient.Options().Proxy)

			log.Infof("addr: %s, service: %s, endpoint: %s, method: %s", addr, req.Service(), req.Endpoint(), req.Method())

			// copy the metadata to prevent race
			md = metadata.Copy(md)
			log.Infof("micro router filter receive md: %v", md)
			if rf, ok := md["X-Micro-Router-Filter"]; ok && len(rf) > 0 {
				// 在有X-Micro-Router-Filter时覆盖Micro-Router
				// 删除已有Micro-Router
				delete(md, "Micro-Router")

				filter := strings.Split(rf, ";")
				for _, v := range filter {
					f := strings.Split(v, ":")
					if len(f) != 2 {
						log.Error("micro router filter format error")
						continue
					}

					if f[0] == req.Service() {
						router := f[1]
						md["Micro-Router"] = router
						break
					}
				}
			}
			log.Infof("micro router filter send md: %v", md)
			ctx = metadata.NewContext(ctx, md)
			return callFunc(ctx, addr, req, rsp, opts)
		}
	}

}
