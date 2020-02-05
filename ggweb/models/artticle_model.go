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

// 存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var articleRowsNum = 0

// 只有首次获取行数的时候采取统计表里的行数
func GetArticleRowsNum () int {
	if articleRowsNum == 0 {
		articleRowsNum = QueryArticleRowNum()
	}
	return articleRowsNum
}

// 查询文章的总条数
func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}

// 还要考虑问题，就是当新增或者删除文章的时候，数据总量会发生变化，所以需要修改下；
func SetArticleRowsNum(){
	articleRowsNum = QueryArticleRowNum()
}

// 定义添加文章；
func AddArticle(article Article)(int64,error){
	i, err := insertArticle(article)
	SetArticleRowsNum()
	return i,err 
}

func insertArticle(article Article)(int64,error){
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