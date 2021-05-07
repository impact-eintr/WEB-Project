package main

import (
	"Zinx/ziface"
	"Zinx/znet"
	"fmt"
)

type PingRouter struct {
	znet.BaseRouter
}

// Ping PreHandle
//func (p *PingRouter) PreHandle(request ziface.IRequest) {
//	fmt.Println("call Router PreHandle")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping\n"))
//	if err != nil {
//		fmt.Println("call back before error")
//	}
//}

// Ping Handle
func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping..ping.\n"))
	if err != nil {
		fmt.Println("call back error")
	}
}

// Ping PostHandle
//func (p *PingRouter) PostHandle(request ziface.IRequest) {
//	fmt.Println("Call Router HostHandle...")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("afer ping\n"))
//	if err != nil {
//		fmt.Println("call back after ping error")
//	}
//}

func main() {
	fmt.Println("Hello Zinx")
	// 1. 创建server句柄
	s := znet.NewServer("[zinx v0.5]")
	// 2. 注册路由a
	s.AddRouter(&PingRouter{})
	// 3. 启动server
	s.Run()
}
