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
	page, _ := this.GetInt("page")
	if page <= 0{
		page = 1
	}
	var artList []models.Article
	artList, _ = models.FindArticleWithPage(page)
	this.Data["PageCode"] = 1
	this.Data["HasFooter"] = true
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