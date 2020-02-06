package controllers

import (
	"fmt"
	"goWeb/ggweb/models"
	"time"
	"log"
)

type AddArticleController struct {
	BaseController
}

type UpdateArticleController struct{
	BaseController
}

type TagsController struct {
	BaseController
}

type DeleteArticleController struct{
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



// 当访问update 的时候触发get方法，相应的页面是通过tplname这个属性指定的返回给客户端的页面
func (this *UpdateArticleController) Get(){
	id,_ := this.GetInt("id")
	fmt.Println("更新查询的id：",id)
	// 先获取id所对应的文章信息；
	art ：= models.QueryArticleWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Tags"] = art.Tags
	this.Data["Short"] = art.Short
	this.Data["Content"] = art.Content
	this.Data["Id"] = art.Id
	this.TplName = "write_article.html"
}


func (this *UpdateArticleController) Post(){
	id,_ := this.GetInt(id)
	fmt.Println("更新的id：",id)
	// 获取浏览器传输的数据，通过表单的name 属性获取值
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	// 实例化model，修改数据库
	art := models.Article{id,title,tags,short,content,"",0}
	_,err := models.UpdateArticle(art)
	// 返回数据给浏览器
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "更新成功"}
	}else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "更新失败"}
	}
	this.ServeJSON()
}

func (this *TagsController) Get() {
	tags := models.QueryArticleWithParam("tags")
	fmt.Println(models.HandleTagsListData(tags))
	this.Data["Tags"] = models.HandleTagsListData(tags)
	this.TplName = "tags.html"
}

func (this *DeleteArticleController) Get(){
	artID,_ := this.GetInt("id")
	fmt.Println("需要删除的id为：",.artID)
	_,err := models.DeleteArticle(artID)
	if err != nil {
		log.Println(err)
	}
	this.Redirect("/",302)
}