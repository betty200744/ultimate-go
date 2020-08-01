package service

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	"go.elastic.co/apm/module/apmgin"
	"gobyexample/awesome-go/gin/handler/publish"
	"net/http"
)

func SetUpRouter() *gin.Engine {
	engine := gin.Default()
	engine.Use(apmgin.Middleware(engine))
	engine.Use(ginhttp.Middleware(opentracing.GlobalTracer(), ginhttp.OperationNameFunc(func(r *http.Request) string {
		return r.URL.Path
	})))
	engine.GET("/publish", publish.Publish)
	return engine
}
