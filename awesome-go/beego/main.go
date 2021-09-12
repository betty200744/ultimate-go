package main

import (
	"github.com/astaxie/beego"
	"ultimate-go/awesome-go/beego/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
