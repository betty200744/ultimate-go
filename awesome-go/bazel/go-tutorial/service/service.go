package service

import (
	"ultimate-go/awesome-go/bazel/go-tutorial/handler/publish"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/", publish.Publish)
	return engine
}

func SetUpRouter1() *gin.Engine {
	engine := gin.Default()
	engine.GET("/", publish.Publish)
	return engine
}
