package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Id   string `uri:"id"`
	Name string `uri:"name"`
}

func main() {
	route := gin.Default()
	pprof.Register(route)
	/*	route.GET("/:name/:id", func(c *gin.Context) {
		person := new(Person)
		if err := c.BindUri(person); err != nil {
			c.JSON(400, gin.H{"msg": err})
		}
		c.JSON(200, gin.H{"id": person.Id, "name": person.Name})
	})*/
	route.Run(":8080")
}
