package controller

import (
	"fmt"

	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	fmt.Println("register controller.....")
	this.TplName = "register.html"
}
