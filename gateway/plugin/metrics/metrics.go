package metrics

import (
	"github.com/micro/micro/plugin"
)

func NewPlugin(opts ...Option) plugin.Plugin {
	return newPrometheus(opts...)
}
