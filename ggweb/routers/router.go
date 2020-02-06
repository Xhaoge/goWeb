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
	// 写文章
	beego.Router("/article/add",&controllers.AddArticleController{})
	// 显示文章内容
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	// 更新文章内容
	beego.Router("/article/update",&controllers.UpdateArticleController{})
	// 删除文章
	beego.Router("/article/delete",&controllers.DeleteArticleController{})
	// 标签
	beego.Router("/tags",&controllers.TagsController{})
}