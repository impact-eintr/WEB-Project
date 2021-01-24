package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/fatih/color"
	"github.com/howeyc/gopass"
)

type Option int

const (
	_ Option = iota
	LOGIN
	LOGUP
	EXIT
)

type User struct {
	Uid string
	Pwd string
}

func LogIn(user User) {
	//登录封装
}

func TalkToServer() {

	conn, err := net.Dial("tcp", "127.0.0.1:6066")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	color.Green("连接成功\n 本地地址%v 远端地址%v\n",
		conn.LocalAddr(),
		conn.RemoteAddr())

	buf := bufio.NewReader(os.Stdin)
	for {
		line, _, _ := buf.ReadLine()
		if string(line) == "exit" {
			log.Println("退出")
			return
		}

		//发送
		_, err := conn.Write(line)
		if err != nil {
			log.Println(err)
		}

		//得到回应
		res := make([]byte, 4096)
		n, err := conn.Read(res)
		if err != nil {
			log.Println(err)
		}
		log.Printf("%v", string(res[:n]))
	}
}

func main() {

	var option Option
	var loop bool
	var user User
	for {
		color.Green("----------------- Golang Chat ------------------")
		color.Green("\t\t\t1 登录\t\t\t")
		color.Green("\t\t\t2 注册\t\t\t")
		color.Green("\t\t\t3 退出\t\t\t")

		fmt.Scanf("%d\n", &option)

		switch option {
		case LOGIN:
			color.Cyan("登录中......\n")
			loop = false
		case LOGUP:
			color.Cyan("跳转中......\n")
			loop = false
		case EXIT:
			color.Cyan("退出中......\n")
			loop = false
		default:
			color.Red("输入有误")
		}

		if option == LOGIN {

			color.Yellow("Input You Uid please......\n")
			fmt.Scanf("%s\n", &user.Uid)
			color.Yellow("Input You Passwd please......\n")
			temp, _ := gopass.GetPasswdMasked()
			user.Pwd = string(temp)

			LogIn(user)

		} else if option == LOGUP {

		}
		fmt.Println(loop)
	}

}
