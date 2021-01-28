package process

import (
	"TCPDemo/server/common"
	"TCPDemo/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *common.Message) {
	for id, up := range userList.onlineUsers {
		go this.SendMesToEachOnlineUser(id, []byte(mes.Data), up.Conn)

	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(uid string, data []byte, conn net.Conn) {

	var mes common.Message

	mes.Type = common.SmsMesType
	mes.Data = string(data)

	data, err := json.Marshal(mes)

	tf := &utils.Transfer{
		Conn: conn,
	}
	fmt.Println(string(data))
	err = tf.PkgWrite(data)
	if err != nil {
		fmt.Println(err)
		//发不出去可能是对端离线 可以设计离线留言功能
	}
}
