package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/openzipkin-contrib/zipkin-go-opentracing"
	"log"
	"net"
	"os"
)

func GetIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func initZipkinConfig(service string) (opentracing.Tracer, zipkintracer.Collector) {
	url := "http://zipkin.meetwhale.com"
	// create collector.
	collector, err := zipkintracer.NewHTTPCollector(url + ":9411/api/v1/spans")
	if err != nil {
		log.Fatal(err)
	}
	ip := GetIp()
	fmt.Println(ip)
	// create recorder.
	recorder := zipkintracer.NewRecorder(collector, true, ip, service)

	// create tracer.
	tracer, err := zipkintracer.NewTracer(
		recorder,
		zipkintracer.ClientServerSameSpan(true),
		zipkintracer.TraceID128Bit(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	// explicitly set our tracer to be the default tracer.
	opentracing.SetGlobalTracer(tracer)
	return tracer, collector
}

func main() {
	tracer, collector := initZipkinConfig("whale-test")
	defer collector.Close()
	span := tracer.StartSpan("say-hello")
	if len(os.Args) != 2 {
		panic("Error: expect two argument")
	}
	who := os.Args[1]

	fmt.Println("hello: ", who)
	span.Finish()
}
