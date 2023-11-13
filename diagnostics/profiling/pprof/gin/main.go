package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	"go.elastic.co/apm/module/apmgin"
	"log"
	"net/http"
	_ "net/http/pprof"
	"ultimate-go/diagnostics/profiling/pprof/handler"
)

func main() {
	engine := gin.Default()
	engine.Use(apmgin.Middleware(engine))
	engine.Use(ginhttp.Middleware(opentracing.GlobalTracer(), ginhttp.OperationNameFunc(func(r *http.Request) string {
		return r.URL.Path
	})))
	engine.GET("/concat/", handler.GinConcatHandler)
	engine.GET("/fib/", handler.GinFibHandler)
	pprof.Register(engine)
	log.Fatal(engine.Run(":6060"))
}
