package process

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"TCPDemo/server/common"
	"TCPDemo/server/module"
	"TCPDemo/server/utils"
)

type UserProcess struct {
	Conn net.Conn
	//表明这个连接是哪个用户的conn
	Uid string
}

func (this *UserProcess) ServerProcessLogin(mes *common.Message) (err error) {
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("处理登陆信息反序列化失败:", err)
		return
	}

	var resMes common.Message
	var loginRes common.LoginRes

	user, err := module.MyUserDao.Login(loginMes.Uid, loginMes.Pwd)
	if err != nil {
		if err == module.ERROR_USER_NOTEXITS {

			loginRes.Code = 404
			loginRes.Error = err.Error()
		} else if err == module.ERROR_USER_PWD {

			loginRes.Code = 403
			loginRes.Error = err.Error()
		} else {

			loginRes.Code = 500
			loginRes.Error = err.Error()
		}

	} else {
		loginRes.Code = 200
		loginRes.Uid = user.Uid
		loginRes.Uname = user.Uname

		LoginCh <- user.Uid
		LoginCh <- user.Uname
		defer close(LoginCh)

		this.Uid = user.Uid
		userList.AddOnlineUser(this)

		log.Println(user.Uid, user.Uname, "上号")
	}

	//对响应数据进行序列化
	data, err := json.Marshal(loginRes)
	if err != nil {
		fmt.Println("json.Marshal failed ", err)
		return
	}

	//封装响应消息
	resMes.Type = common.LoginResType
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal failed ", err)
		return
	}
	//发送
	tf := utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.PkgWrite(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	return

}

func (this *UserProcess) ServerProcessRegister(mes *common.Message) (err error) {
	var registerMes common.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("处理登陆信息反序列化失败:", err)
		return
	}

	var resMes common.Message
	var registerRes common.RegisterRes

	err = module.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == module.ERROR_USER_EXITS {

			registerRes.Code = 400
			registerRes.Error = err.Error()
		}

	} else {
		registerRes.Code = 200
		registerRes.Uid = registerMes.User.Uid
		registerRes.Uname = registerMes.User.Uname

		RegisterCh <- registerRes.Uid
		RegisterCh <- registerRes.Uname
		fmt.Println(registerRes.Uid, registerRes.Uname)
	}

	//对响应数据进行序列化
	data, err := json.Marshal(registerRes)
	if err != nil {
		fmt.Println("json.Marshal failed ", err)
		return
	}

	//封装响应消息
	resMes.Type = common.LoginResType
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal failed ", err)
		return
	}
	//发送
	tf := utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.PkgWrite(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	return

}
