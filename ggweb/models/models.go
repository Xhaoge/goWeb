package models

import (
	"fmt"
	"goWeb/ggweb/utils"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int //0 正常状态， 1 删除
	Createtime int64
}

// 数据库操作
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username,password,status,createtime) value (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

/*
按条件查询
补充：scan函数会识别空格左右的内容，哪怕换行符号存在也不会影响scan对内容的获取。
scanln函数会识别空格左右的内容，但是一旦遇到换行符就会立即结束，不论后续还是否存在需要带输入的内容。
*/

// 按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	fmt.Println("row", row)
	id := 0
	row.Scan(&id)
	fmt.Println("row2", &row)
	fmt.Println("row2", id)
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

//根据用户名和密码，查询id
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}
