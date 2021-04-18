package logic

import (
	"bluebell/dao/mysql"
	sf "bluebell/pkg/snowflake"
)

func SignUp() {
	// 检查用户是否已经注册
	mysql.QueryUserByName()

	// 生成UID
	sf.GenID()

	// 存入数据库
	mysql.InsertUser()

}
