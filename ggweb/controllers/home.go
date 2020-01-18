package controllers

import (
	"fmt"

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
	fmt.Println("IsLogin:",this.IsLogin,this.Loginuser)
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