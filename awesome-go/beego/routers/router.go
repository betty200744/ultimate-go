package routers

import (
	"github.com/astaxie/beego"
	"ultimate-go/awesome-go/beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
