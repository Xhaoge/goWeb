package controllers

import (
	"fmt"
	"time"
	"github.com/astaxie/beego"
	"goWeb/ggweb/models"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	fmt.Println("register controller. get....")
	this.TplName = "register.html"
}

func (this *RegisterController) Post() {
	fmt.Println("register controller. post....")
	// 获取表单信息
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	fmt.Println(username,password)

	// 首先检查下是否已经在注册过了，如果已经注册，返回错误；
	id := models.QueryUserWithUsername(username)
	fmt.Println("注册时检查id：",id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code":0,"message":"用户名已经存在，请重新输入；"}
		this.ServeJSON()
		return
	}

	// 检查下 俩次密码 是否输入一致
	if password != repassword {
		this.Data["json"] = map[string]interface{}{"code":0,"message":"密码输入不一致，请重新输入"}
		this.ServeJSON()
		return
	}

	// 注册用户名和密码 todo 加密入库
	user := models.User{0,username,password,0,time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"注册失败；"}
	}else {
		this.Data["json"] = map[string]interface{}{"code":0,"message":"注册成功；"}
	}
	this.ServeJSON()
}