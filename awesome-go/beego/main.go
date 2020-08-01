package main

import (
	"github.com/astaxie/beego"
	"gobyexample/awesome-go/beego/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
