package models

import (
	"fmt"
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
