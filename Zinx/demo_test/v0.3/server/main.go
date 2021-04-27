package main

import (
	"Zinx/znet"
	"fmt"
)

func main() {
	fmt.Println("Hello Zinx")
	// 1. 创建server句柄
	s := znet.NewServer("[zinx v0.1]")
	// 2. 启动server
	s.Run()
}
