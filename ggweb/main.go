package main

import (
	_ "goWeb/ggweb/routers"
	"goWeb/ggweb/utils"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Info("this is ggweb")
	utils.InitMysql()
	beego.Run()
}
