package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 直接连接远程服务器 得到一个conn连接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("write conn err:", err)
		return
	}

	for {
		_, err := conn.Write([]byte("Zinx服务器你好！！！"))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}

		buf := make([]byte, 512)
		cbt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err", err)
			return
		}

		fmt.Printf("server call back:%s[%d]\n", string(buf), cbt)
		time.Sleep(time.Second)

	}
}
