package main

import (
	"gobyexample/awesome-go/bazel/go-tutorial/service"
)

func main() {
	engine := service.SetUpRouter()
	engine.Run(":8082")
}
