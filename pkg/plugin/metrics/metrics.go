package metrics

import (
	"github.com/stack-labs/starter-kit/pkg/gateway/plugin"
)

//NewPlugin of metrics
func NewPlugin(opts ...Option) plugin.Plugin {
	return newPrometheus(opts...)
}
