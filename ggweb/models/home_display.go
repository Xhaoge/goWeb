package models

import (
	"bytes"
	"fmt"
	"goWeb/ggweb/utils"
	"html/template"
	"strconv"
)

type HomeBlockParam struct {
	Id 			int
	Title 		string
	Tags 		string
	Short 		string
	Content 	string
	Author 		string
	CreateTime 	string
	// 查看文章地址；
	Link 		string
	// 修改文章地址；
	UpdateLink 	string
	DeleteLink 	string
	// 记录是否登录
	IsLogin 	bool
}

// 标签链接
type TagLink struct {
	TagName 	string
	TagUrl 		string
}

func MakeHomeBlocks(articles []Article,isLogin bool) template.HTML{
	fmt.Println("this is makehomeblocks")
	htmlHome := ""
	for _, art := range articles{
		// 将数据中的数据 转换为 首页模板所需要的model
		homeParam := HomeBlockParam{}
		homeParam.Id = art.Id
		homeParam.Title = art.Title
		homeParam.Author = art.Author
		homeParam.Tags = art.Tags
		fmt.Println("---------->tags",art.Tags)
		homeParam.Short = art.Short
		homeParam.Content = art.Content
		homeParam.CreateTime = utils.SwitchTimeStampToData(art.Createtime)
		homeParam.Link = "/article/" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homeParam.IsLogin = isLogin

		//处理变量，parsefile 解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}
		// 就是将html 文件里面的数据替换；
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()

	}
	fmt.Println("htmlhome------>",htmlHome)
	return template.HTML(htmlHome)
}
