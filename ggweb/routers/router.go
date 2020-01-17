package routers

import (
	"github.com/astaxie/beego"
	"goWeb/ggweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register",&controllers.RegisterController{})
}