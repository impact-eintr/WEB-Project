package process

import (
	"github.com/fatih/color"
	"net"
)

func talkToServer() (net.Conn, error) {

	conn, err := net.Dial("tcp", "127.0.0.1:6066")
	if err != nil {
		defer conn.Close()
		return nil, err
	}

	color.Green("连接成功\n 本地地址%v 远端地址%v\n",
		conn.LocalAddr(),
		conn.RemoteAddr())
	return conn, nil
}
