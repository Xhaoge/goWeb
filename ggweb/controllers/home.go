package controllers

import (
	"fmt"
	"goWeb/ggweb/models"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
	Loginuser	interface{}
}

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	tag := this.GetString("tag")
	fmt.Println("首页tag：",tag)
	page, err := this.GetInt("page")
	var artList []models.Article
	if len(tag) > 0 {
		// 按照指定的标签搜索
		artList, _ = models.QueryArticleWithTag(tag)
		this.Data["HasFooter"] = false
	}else {
		if page <= 0 || err != nil {
			page = 1
		}
		artList,_ = models.FindArticleWithPage(page)
		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		this.Data["HasFooter"] = true
	}
	fmt.Println("IsLogin:",this.IsLogin,this.Loginuser)
	this.Data["content"] = models.MakeHomeBlocks(artList,this.IsLogin)
	this.TplName = "home.html"
}

// 判断是否登录
func (this *BaseController) Prepase(){
	loginuser := this.GetSession("loginuser")
	fmt.Println("loginuser cookie----->",loginuser)
	if loginuser != nil {
		this.IsLogin = true
		this.Loginuser = loginuser
	} else {
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}