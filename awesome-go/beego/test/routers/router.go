package routers

import (
	"github.com/astaxie/beego"
	"gobyexample/awesome-go/beego/test/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
