package jaeger

import (
	"fmt"
	"github.com/fajarardiyanto/flt-go-tracer/interfaces"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

type Jaeger struct {
	name   string
	config interfaces.JaegerConfig
}

func NewJaeger(name string, config interfaces.JaegerConfig) interfaces.Jaeger {
	return &Jaeger{
		name:   name,
		config: config,
	}
}

func (j *Jaeger) InitTracer() (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: j.name,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           j.config.LogSpan,
			LocalAgentHostPort: fmt.Sprintf(":%s", j.config.Endpoint),
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
