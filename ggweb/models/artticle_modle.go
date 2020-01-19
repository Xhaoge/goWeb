package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"goWeb/ggweb/utils"
)

type Article struct {
	Id 			int
	Title 		string
	Author  	string
	Tags 		string
	Short 		string
	Content 	string
	Createtime  int64
	// status int 0 为正常，1 为删除  2 为冻结
}

// 定义添加文章；
func AddArticle(article Article)(int64,error){
	fmt.Println("insert into mysql article,article:",article)
	return utils.ModifyDB("insert into article(title,author,tags,short,content,createtime)" +
		"value(?,?,?,?,?,?)",article.Title,article.Author,article.Tags,article.Short,article.Content,article.Createtime)
}

// 定义查询文章；
func FindArticleWithPage(page int)([]Article,error){
	// 从配置中获取每页文章数量；
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("-----------page:",page)
	return QueryArticleWithPage(page,num)
}

func QueryArticleWithPage(page,num int)([]Article,error){
	sql := fmt.Sprintf("limit %d,%d",page*num,num)
	return QueryArticlesWithCon(sql)
}

func QueryArticlesWithCon(sql string)([]Article,error){
	sql = "select id,title,tags,short,content,author,createtime from article "+sql
	rows,err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next(){
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id,&title,&author,&tags,&short,&content,&createtime)
		art := Article{id,title,author,tags,short,content,createtime}
		artList = append(artList,art)
	}
	return artList,nil
}