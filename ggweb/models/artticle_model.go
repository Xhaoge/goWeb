package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"goWeb/ggweb/utils"
	"log"
	"strings"
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

//更新文章
func UpdateArticle(art Article)(int64,error){
	return utils.ModifyDB("select article set title=?,tags=?,short=?,content=? where id=?",art.Title,art.Tags,art.Short,art.Content,art.Id)
}

// 删除文章
func DeleteArticle(artID Int) (int64,error){
	i,err := deleteArticleWithArtId(artID)
	SetArticleRowsNum()
	return i,err
}

func deleteArticleWithArtId(artID int) (int64,error){
	return utils.ModifyDB("delete from article where id=?",artID)
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

func QueryArticleWithId(id int){
	row := utils.QueryRowDB("select id,title,tags,short,content,author,createtime from article where id="+ strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}

//查询标签，返回一个字段的列表
func QueryArticleWithParam(param string) []string {
	rows,_ := utils.QueryDB(fmt.Sprintf("select %s from article",param))
	if err != nil {
		log.Println(err)
	}
	var paramList []string
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList,arg)
	}
	return paramList
}


func HandleTagsListData(tags []string) map[string]int {
	var tagsMap = make(map[string]int)
	for _, tag := range tags {
		taglist := strings.Split(tag,"&")
		for _, value := range taglist {
			tagsMap[value]++
		}
	}
}