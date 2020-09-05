package chain

import (
	"github.com/micro/go-micro/v3/registry"
	"github.com/micro/micro/v3/service/gateway/router"
)

func filterChain(labelKey string, chains []string) router.ServiceFilter {
	return func(old []*registry.Service) []*registry.Service {
		var services []*registry.Service

		chain := ""
		idx := len(chains)
		for _, service := range old {
			serv := new(registry.Service)
			var nodes []*registry.Node

			for _, node := range service.Nodes {
				if node.Metadata == nil {
					continue
				}

				val := node.Metadata[labelKey]
				if len(val) == 0 {
					continue
				}

				if len(chain) > 0 && idx == 0 {
					if chain == val {
						nodes = append(nodes, node)
					}
					continue
				}

				// chains按顺序优先匹配
				ok, i := inArray(val, chains)
				if ok && idx > i {
					// 出现优先链路，services清空，nodes清空
					idx = i
					chain = val
					services = services[:0]
					nodes = nodes[:0]
				}

				if ok && idx == i {
					nodes = append(nodes, node)
				}
			}

			// only add service if there's some nodes
			if len(nodes) > 0 {
				// copy
				*serv = *service
				serv.Nodes = nodes
				services = append(services, serv)
			}
		}

		if len(services) == 0 {
			return old
		}

		return services
	}
}

func inArray(s string, d []string) (bool, int) {
	for k, v := range d {
		if s == v {
			return true, k
		}
	}
	return false, 0
}
