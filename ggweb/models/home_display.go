package models

import (
	"bytes"
	"fmt"
	"html/template"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type HomeBlockParam struct {
	Id 			int
	Title 		string
	Tags 		[] TagLink
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

// 分页的结构体
type HomeFooterPageCode struct {
	HasPre		bool
	HasNext 	bool
	ShowPage 	string
	PreLink 	string
	NextLink 	string
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
		homeParam.Tags = createTagsLinks(art.Tags)
		fmt.Println("---------->tags",art.Tags)
		homeParam.Short = art.Short
		homeParam.Content = art.Content
		homeParam.CreateTime =strconv.FormatInt(art.Createtime,10)
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

// 将tags 字符串转化成首页模板所需要的 数据结构；
func createTagsLinks(tags string) []TagLink {
	var tagLink []TagLink
	tagsPamar := strings.Split(tags, "&")
	for _, tag := range tagsPamar {
		tagLink = append(tagLink, TagLink{tag, "/?tag="+tag})
	}
	return tagLink
}

// 查询文章和修改方法
func ConfigHomeFooterPageCode(page int) HomeFooterPageCode{
	pageCode := HomeFooterPageCode{}
	// 查询出总的条数
	num := GetArticleRowsNum()
	// 从配置文件中读取每页显示的条数
	pageRow, _ := beego.AppConfig.Int("articleListPageNum")
	//计算出总页数
	allPageNum := (num-1)/pageRow +1
	pageCode.ShowPage = fmt.Sprintf("%d/%d",page,allPageNum)
	// 当前页数大于总页数，那么下一页的按钮不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	pageCode.PreLink = "/?page="+ strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode
}
