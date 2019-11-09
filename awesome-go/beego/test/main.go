package main

import (
	"github.com/astaxie/beego"
	"gobyexample/awesome-go/beego/test/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
