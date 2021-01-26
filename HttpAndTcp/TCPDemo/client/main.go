package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/howeyc/gopass"

	"TCPDemo/client/common"
	"TCPDemo/client/process"
)

func main() {
	var option common.Option
	var user common.User
	for {
		color.Green("----------------- Golang Chat ------------------")
		color.Green("\t\t1 登录\t\t\t")
		color.Green("\t\t2 注册\t\t\t")
		color.Green("\t\t3 退出\t\t\t")

		fmt.Scanf("%d\n", &option)

		switch option {
		case common.LOGIN:
			color.Yellow("Input You Uid please......\n")
			fmt.Scanf("%s\n", &user.Uid)
			color.Yellow("Input You Passwd please......\n")
			temp, _ := gopass.GetPasswdMasked()
			user.Pwd = string(temp)

			color.Cyan("登录中......\n")
			up := process.UserProcess{}
			up.LogIn(user)

		case common.SIGNIN:
			color.Green("Input You Uid please......\n")
			fmt.Scanf("%s\n", &user.Uid)
			for {
				color.Yellow("Input You Passwd please......\n")
				temp, _ := gopass.GetPasswdMasked()
				color.Yellow("Input You Passwd please again......\n")
				pwd, _ := gopass.GetPasswdMasked()
				if string(temp) != string(pwd) {
					color.Red("两次输入密码不同，重来一次吧...\n")
				} else {
					user.Pwd = string(temp)
					break
				}
			}

			color.Cyan("跳转中......\n")
			up := process.UserProcess{}
			up.Register(user)

		case common.EXIT:
			color.Cyan("退出中......\n")
			os.Exit(0)
		default:
			color.Red("输入有误")
		}

	}

}
