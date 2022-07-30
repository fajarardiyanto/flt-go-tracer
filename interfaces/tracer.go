package interfaces

import (
	"github.com/opentracing/opentracing-go"
	"io"
)

type Tracing interface {
	LoadJaeger(name string, config JaegerConfig) Jaeger
}

type JaegerConfig struct {
	Enable  bool
	Url     string
	Host    string
	Port    string
	LogSpan bool
}

type Jaeger interface {
	InitTracer() (opentracing.Tracer, io.Closer)
}
