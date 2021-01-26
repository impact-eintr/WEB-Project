package module

import (
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

func (this *UserDao) getUserById(conn redis.Conn, uid string) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", uid))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXITS
		}
		return
	}

	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal failed :", err)
		return
	}
	return

}

//登陆校验
func (this *UserDao) Login(uid, pwd string) (user *User, err error) {
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
