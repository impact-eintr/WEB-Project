package module

import (
	"TCPDemo/server/common"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//构造函数
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn, uid string) (user *common.User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", uid))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXITS
		}
		return
	}

	user = &common.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal failed :", err)
		return
	}
	return

}

//登陆校验
func (this *UserDao) Login(uid, pwd string) (user *common.User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserById(conn, uid)
	if err != nil {
		return
	}

	//用户存在
	if user.Pwd != pwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *common.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.Uid)
	if err == nil {
		err = ERROR_USER_EXITS
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = conn.Do("HSet", "users", user.Uid, string(data))
	if err != nil {
		fmt.Println("数据库内部错误：", err)
		return
	}
	return
}
