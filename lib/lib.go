package lib

import (
	"github.com/fajarardiyanto/flt-go-tracer/interfaces"
	"github.com/fajarardiyanto/flt-go-tracer/lib/jaeger"
)

type Modules struct {
}

func NewLib() interfaces.Tracing {
	return &Modules{}
}

func (*Modules) LoadJaeger(name string, config interfaces.JaegerConfig) interfaces.Jaeger {
	return jaeger.NewJaeger(name, config)
}
