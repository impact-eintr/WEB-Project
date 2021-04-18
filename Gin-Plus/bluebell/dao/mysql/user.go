package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

func CheckUserExist(username string) error {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	err := db.Get(&count, sqlStr, username)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("用户已经存在")
	}
	return nil
}

func InsertUser(user *models.User) (err error) {
	// 密码加密
	user.Password = encryptPassword(user.Password)

	// 执行SQL语句入库
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return

}

const salt string = `impact-eintr`

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
