package main

import (
	_ "goWeb/ggweb/routers"
	"goWeb/ggweb/utils"

	"github.com/astaxie/beego"
)

func main() {
	beego.Info("this is ggweb")
	utils.InitMysql()
	beego.Run()
}
