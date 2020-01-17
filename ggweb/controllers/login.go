package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"goWeb/ggweb/models"
)

type LoginController struct{
	beego.Controller
}


// 定义login 页面get请求；

func (this *LoginController) Get(){
	fmt.Println("this is LoginController get .......")
	this.TplName = "login.html"
}

// 定义login 页面post登录请求；
func (this *LoginController) Post(){
	fmt.Println("this is LoginController post ......")
	// 获取登录信息
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:", username, "password:", password)
	// 从数据库中读取用户信息；
	id := models.QueryUserWithParam(username,password)
	fmt.Println("login 登录时查询到的Id号为：",id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code":1,"message":"登录成功"}
	} else{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"登录失败"}
	}
	this.ServeJSON()

}