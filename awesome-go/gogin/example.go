package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		data := map[string]interface{}{
			"a": "a",
			"b": "b",
		}
		c.JSON(200, gin.H{"message": "pong"})
		c.AsciiJSON(200, data)
	})
	r.Run(":3003")

}
