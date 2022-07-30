package jaeger

import (
	"context"
	"fmt"
	"github.com/fajarardiyanto/flt-go-tracer/interfaces"
	"github.com/fajarardiyanto/flt-go-utils/caller"
	"github.com/fajarardiyanto/flt-go-utils/parser"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"runtime"
)

func RootSpan(ctx context.Context) (opentracing.Span, opentracing.Tracer) {
	parentSpan := opentracing.SpanFromContext(ctx)
	tracer := parentSpan.Tracer()

	// Get caller function name, file and line
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	parentSpan.SetTag("caller", callerDetails)

	return parentSpan, tracer
}

func CreateRootSpan(ctx context.Context, name string) (opentracing.Span, opentracing.Tracer) {
	parentSpan, _ := opentracing.StartSpanFromContext(ctx, name)
	tracer := parentSpan.Tracer()
	parentSpan.SetTag("name", name)

	frame, _ := caller.GetCaller(2)
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	parentSpan.SetTag("caller", callerDetails)

	return parentSpan, tracer
}

func CreateSubChildSpan(parentSpan opentracing.Span, name string) opentracing.Span {
	sp := opentracing.StartSpan(
		name,
		opentracing.ChildOf(parentSpan.Context()))
	sp.SetTag("name", name)

	frame, _ := caller.GetCaller(2)
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	sp.SetTag("caller", callerDetails)

	return sp
}

func LogRequest(sp opentracing.Span, req interface{}) {
	sp.LogFields(log.Object(string(interfaces.Request), parser.Stringify(req)))
}

func LogObject(sp opentracing.Span, name string, resp interface{}) {
	sp.LogFields(log.Object(name, parser.Stringify(resp)))
}

func LogResponse(sp opentracing.Span, resp interface{}) {
	sp.LogFields(log.Object(string(interfaces.Response), parser.Stringify(resp)))
}

func LogError(sp opentracing.Span, err error) {
	sp.SetTag("error", true)
	sp.LogFields(log.Object(string(interfaces.Error), err))
}
