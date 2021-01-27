package process

import (
	"TCPDemo/client/common"
	"TCPDemo/client/utils"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"log"
)

type UserProcess struct {
}

//登录封装
func (this *UserProcess) LogIn(user common.User) {

	conn, err := connToServer()
	if err != nil {
		log.Println(err)
		return
	}

	//开始登陆
	var mes common.Message

	//构建登陆信息
	var loginMes common.LoginMes
	loginMes.Uid = user.Uid
	loginMes.Pwd = user.Pwd

	//协议数据部分序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		log.Println(err)
		return
	}
	mes.Type = common.LoginMesType
	mes.Data = string(data)

	//协议序列化
	data, err = json.Marshal(mes)
	if err != nil {
		log.Println(err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	tf.PkgWrite(data)

	resp, err := tf.PkgRead()
	var temp common.LoginRes
	json.Unmarshal([]byte(resp.Data), &temp)

	if temp.Code == 200 {
		uname := color.CyanString(" %s ", temp.Uname)
		fmt.Printf("%s欢迎回来\n", uname)

		go talkToServer(conn)

		for {
			ShowMenu()
		}
	} else {
		color.Red("认证失败 %s\n", temp.Error)
	}

}

func (this *UserProcess) Register(user common.User) {

	conn, err := connToServer()
	if err != nil {
		log.Println(err)
		return
	}

	//开始登陆
	var mes common.Message

	//构建登陆信息
	var registerMes common.RegisterMes
	registerMes.User.Uid = user.Uid
	registerMes.User.Pwd = user.Pwd
	registerMes.User.Uname = user.Uname

	//协议数据部分序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		log.Println(err)
		return
	}
	mes.Type = common.RegisterMesType
	mes.Data = string(data)

	//协议序列化
	data, err = json.Marshal(mes)
	if err != nil {
		log.Println(err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	tf.PkgWrite(data)

	resp, err := tf.PkgRead()
	log.Println("test")
	var temp common.RegisterRes
	json.Unmarshal([]byte(resp.Data), &temp)

	if temp.Code == 200 {
		uname := color.CyanString(" %s ", temp.Uname)
		fmt.Printf("%s，注册成功，可以上号了\n", uname)

	} else {
		color.Red("注册失败 %s\n", temp.Error)
	}
}
