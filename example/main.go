package main

import (
	"context"
	"github.com/fajarardiyanto/flt-go-tracer/interfaces"
	"github.com/fajarardiyanto/flt-go-tracer/lib"
	jaeger2 "github.com/fajarardiyanto/flt-go-tracer/lib/jaeger"
)

func main() {
	jaeger := lib.NewLib()
	_, closer := jaeger.LoadJaeger("Module", interfaces.JaegerConfig{
		Host:    "0.0.0.0",
		Port:    "6831",
		LogSpan: true,
	}).InitTracer()

	defer closer.Close()

	span, _ := jaeger2.CreateRootSpan(context.Background(), "TESTING")
	defer span.Finish()

	sp := jaeger2.CreateSubChildSpan(span, "HELLO")
	defer sp.Finish()

	jaeger2.LogRequest(sp, "Example Module")
}
