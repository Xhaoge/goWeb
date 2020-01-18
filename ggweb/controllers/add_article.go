package controllers

import (
	"fmt"
	"goWeb/ggweb/models"
	"time"
)

type AddArticleController struct {
	BaseController
}

/*
	当访问/add 路径的时候触发AddArticleController的get方法，响应的页面但是通过tplname
*/

func (this *AddArticleController) Get(){
	this.TplName = "write_article.html"
}

// 添加文章逻辑实现；
func (this *AddArticleController) Post(){
	fmt.Println("this AddArticleController post....")
	// 获取浏览器传输的数据，通过表单的name 属性获取值；
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	fmt.Printf("输入文章为：",title,tags)

	// 实例化model，将他的数据存入到数据库中；
	art := models.Article{0,title,"xhaoge",tags,short,content,time.Now().Unix()}
	_, err := models.AddArticle(art)
	fmt.Println("insert article err :",err)

	// 返回数据给浏览器；
	var response map[string]interface{}
	if err != nil {
		response = map[string]interface{}{"code":1,"message":"ok"}
	} else {
		response = map[string]interface{}{"code":0,"message":"error"}
	}
	this.Data["json"] = response
	this.ServeJSON()
}
