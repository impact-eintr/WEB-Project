package process

import (
	"encoding/json"
	"fmt"
	"net"

	"TCPDemo/server/common"
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

	if loginMes.Uid == "111" && loginMes.Pwd == "hhh" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "用户不存在"
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

func (this *UserProcess) ServerProcessRegister() (err error) {
	return
}
