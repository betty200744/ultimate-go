package main

import (
	"fmt"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"net/http"
)

// jaeger:  http://localhost:16686/

func main() {

	/*
		Setting up your tracer jaeger tracer
	*/
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		ServiceName: "your_service_name",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, _ := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	someFunction()
}

func someFunction() {
	/*
		Start Trace With Parent Span
	*/
	globalTrace := opentracing.GlobalTracer()

	parent := globalTrace.StartSpan("hello 2")
	parent.SetTag("tag1", "value cc")
	parent.SetBaggageItem("Baggage1", "value bagg")
	defer parent.Finish()

	/*
		Create a Child Span
	*/
	clientSpan := globalTrace.StartSpan(
		"world", opentracing.ChildOf(parent.Context()))

	/*
		Make an HTTP request
	*/
	url := "http://localhost:8082/publish"
	req, _ := http.NewRequest("GET", url, nil)

	// Set some tags on the clientSpan to annotate that it's the client span. The additional HTTP tags are useful for debugging purposes.

	// Inject the client span context into the headers
	globalTrace.Inject(clientSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	resp, _ := http.DefaultClient.Do(req)

	fmt.Println(resp.Status)
	defer clientSpan.Finish()
}

func HttpFunction() {

}
