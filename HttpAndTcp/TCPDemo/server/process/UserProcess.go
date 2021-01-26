package process

import (
	"encoding/json"
	"fmt"
	"net"

	"TCPDemo/server/common"
	"TCPDemo/server/module"
	"TCPDemo/server/utils"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerProcessLogin(mes *common.Message) (err error) {
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("处理登陆信息反序列化失败:", err)
		return
	}

	var resMes common.Message
	var loginResMes common.LoginRes

	user, err := module.MyUserDao.Login(loginMes.Uid, loginMes.Pwd)
	if err != nil {
		if err == module.ERROR_USER_NOTEXITS {

			loginResMes.Code = 404
			loginResMes.Error = err.Error()
		} else if err == module.ERROR_USER_PWD {

			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {

			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		}

	} else {
		loginResMes.Code = 200
		loginResMes.Uid = user.Uid
		LoginCh <- user.Uid
		loginResMes.Uname = user.Uname
		LoginCh <- user.Uname
		fmt.Println(user.Uid, user.Uname, "上号")
	}

	//对响应数据进行序列化
	data, err := json.Marshal(loginResMes)
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

	user, err := module.MyUserDao.Register(registerMes)
	if err != nil {
		if err == module.ERROR_USER_EXITS {

			registerRes.Code = 400
			registerRes.Error = err.Error()
		}

	} else {
		registerRes.Code = 200
		RegisterCh <- registerRes.Uid
		RegisterCh <- registerRes.Uname
	}

	//对响应数据进行序列化
	data, err := json.Marshal(loginResMes)
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
