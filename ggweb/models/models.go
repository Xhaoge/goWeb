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
	return utils.ModifyDB("insert into users(username,password,status,createtimr) value (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

/*
按条件查询
补充：scan函数会识别空格左右的内容，哪怕换行符号存在也不会影响scan对内容的获取。
scanln函数会识别空格左右的内容，但是一旦遇到换行符就会立即结束，不论后续还是否存在需要带输入的内容。
*/

func QueryUserWightCon(con string) int {
	sql := fmt.Sprint("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

// 根据用户名查询ID
func QueryUserWightParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}
