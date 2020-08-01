package main

import (
	"gobyexample/awesome-go/gin/service"
)

func main() {
	engine := service.SetUpRouter()
	engine.Run(":8082")
}
