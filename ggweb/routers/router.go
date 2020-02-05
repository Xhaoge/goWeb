package routers

import (
	"github.com/astaxie/beego"
	"goWeb/ggweb/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register",&controllers.RegisterController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/exit",&controllers.ExitController{})
	beego.Router("/article/add",&controllers.AddArticleController{})
	// 显示文章内容
	beego.Router("/article/:id", &controllers.ShowArticleController{})
}