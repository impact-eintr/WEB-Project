package process

import (
	"TCPDemo/client/utils"
	"fmt"
	"github.com/fatih/color"
	"log"
	"net"
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
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.PkgRead()
		if err != nil {
			log.Println("tf.ReadPkg err=", err)
			return
		}
		fmt.Printf("%v\n", mes)
	}
}
