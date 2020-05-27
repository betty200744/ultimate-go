package main

import (
	"fmt"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
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
		ServiceName: "hello2",
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
	// tags key:value for query, filter and comprehend trace
	parent.SetTag("tag1", "t1")
	parent.SetTag("tag2", "t2")
	// logs key:value for log, debug
	parent.LogFields(
		log.String("event", "soft error"),
		log.String("type", "cache timeout"),
		log.Int("waited.millis", 1500))
	defer parent.Finish()

	/*
			Create a Child Span
		    Make an HTTP request
	*/
	url := "http://localhost:8082/publish"
	clientSpan := globalTrace.StartSpan(
		"world", opentracing.ChildOf(parent.Context()))
	clientSpan.SetBaggageItem("Baggage_url", url)
	clientSpan.SetBaggageItem("Baggage_b2", "b2")
	clientSpan.SetTag("tag3", "t3")
	clientSpan.SetTag("tag4", "t4")
	// logs key:value for log, debug
	clientSpan.LogFields(
		log.String("url", url),
		log.String("log2", "l2"),
		log.Int("waited.millis", 1500))
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
