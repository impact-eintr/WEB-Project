package process

import (
	"TCPDemo/client/common"
	"TCPDemo/client/utils"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/fatih/color"
)

func connToServer() (net.Conn, error) {

	conn, err := net.Dial("tcp", "127.0.0.1:6066")
	if err != nil {
		defer conn.Close()
		return nil, err
	}

	color.Green("连接成功\n")
	return conn, nil
}

func talkToServer(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("==========================")
		mes, err := tf.PkgRead()
		if err != nil {
			log.Println("tf.ReadPkg err=", err)
			return
		}

		switch mes.Type {
		case common.NotifyUserStatusMesType:
			//更新客户端维护的usermap
			var notifyUserStatusMes common.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		default:
			fmt.Println("消息类型未知")
		}

	}
}
