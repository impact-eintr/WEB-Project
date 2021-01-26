package process

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

type option int

const (
	_ option = iota
	ONLINE
	SEND
	INFO
	EXIT
)

func ShowMenu() {
	color.Yellow("=========1 显示在线用户列表========")
	color.Yellow("=========2 发送消息================")
	color.Yellow("=========3 信息列表================")
	color.Yellow("=========4 退出系统================")

	var op option
	fmt.Scanf("%d\n", &op)
	switch op {
	case ONLINE:
		fmt.Println("显示用户列表")
	case SEND:
		fmt.Println("发送消息")
	case INFO:
		fmt.Println("消息列表")
	case EXIT:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入有误")
	}
}
