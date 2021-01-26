package process

import (
	"TCPDemo/client/common"
	"TCPDemo/client/utils"
	"encoding/json"
	"log"
)

type UserProcess struct {
}

//登录封装
func (this *UserProcess) LogIn(user common.User) {

	conn, err := talkToServer()
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
	log.Println(resp)
	var temp common.LoginRes
	json.Unmarshal([]byte(resp.Data), &temp)
	log.Println(temp.Code)
}
