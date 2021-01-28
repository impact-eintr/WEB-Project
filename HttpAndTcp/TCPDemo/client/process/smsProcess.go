package process

import (
	"TCPDemo/client/common"
	"TCPDemo/client/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes common.Message
	mes.Type = common.SmsMesType

	var smsMes common.SmsMes
	smsMes.Content = content
	smsMes.Uid = CurUser.Uid
	smsMes.Ustatus = CurUser.Ustatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println(err)
		return err
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(err)
		return err
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.PkgWrite(data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}
