package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	"go.elastic.co/apm/module/apmgin"
	"net/http"
)

type Person struct {
	Id   string `uri:"id"`
	Name string `uri:"name"`
}

func main() {
	engine := gin.Default()
	engine.Use(apmgin.Middleware(engine))
	engine.Use(ginhttp.Middleware(opentracing.GlobalTracer(), ginhttp.OperationNameFunc(func(r *http.Request) string {
		return r.URL.Path
	})))
	engine.GET("/publish", func(c *gin.Context) {
		person := new(Person)
		person.Id = "1"
		person.Name = "betty"
		c.JSON(200, gin.H{"id": person.Id, "name": person.Name})
	})
	engine.Run(":8082")
}
