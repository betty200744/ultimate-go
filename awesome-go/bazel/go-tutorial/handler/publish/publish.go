package publish

import "github.com/gin-gonic/gin"

type Person struct {
	Id   string `uri:"id"`
	Name string `uri:"name"`
}

func Publish(c *gin.Context) {
	person := new(Person)
	person.Id = "1"
	person.Name = "betty"
	c.JSON(200, gin.H{"id": person.Id, "name": person.Name})
}
