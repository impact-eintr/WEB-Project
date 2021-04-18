package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	sf "bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) error {
	// 检查用户是否已经注册
	err := mysql.CheckUserExist(p.UserName)
	if err != nil {
		return err
	}

	// 生成UID
	userID := sf.GenID()
	// 构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.UserName,
		Password: p.Password,
	}

	// 存入数据库
	return mysql.InsertUser(user)

}
