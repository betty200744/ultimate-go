package main

import (
	"fmt"
	"github.com/opentracing/basictracer-go"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"os"
)

// initJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

func main() {
	tracer, closer := initJaeger("hello-world2")
	defer closer.Close()
	span := tracer.StartSpan("say-hello")
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}
	sc, _ := span.Context().(basictracer.SpanContext)
	fmt.Println("this is span traceid: ", sc.TraceID)
	who := os.Args[1]
	println("hello, ", who)
	span.Finish()
}
