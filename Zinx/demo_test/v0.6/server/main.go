package main

import (
	"Zinx/ziface"
	"Zinx/znet"
	"fmt"
)

type PingRouter struct {
	znet.BaseRouter
}

type HelloRouter struct {
	znet.BaseRouter
}

const (
	MsgPing  uint32 = 0
	MsgHello uint32 = 1
)

// Ping Handle
func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	fmt.Println("recv from client:msgID = ", request.GetMsgID(),
		",data = ", string(request.GetMsgData()))

	// 先读取客户端数据 再写回数据到客户端
	err := request.GetConnection().SendMsg(222, []byte("pong...pong..pong.\n"))
	if err != nil {
		fmt.Println("call back error")
	}
}

func (h *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	fmt.Println("recv from client:msgID = ", request.GetMsgID(),
		",data = ", string(request.GetMsgData()))

	// 先读取客户端数据 再写回数据到客户端
	err := request.GetConnection().SendMsg(666, []byte("hello...hello..hello.\n"))
	if err != nil {
		fmt.Println("call back error")
	}
}

func main() {
	fmt.Println("Hello Zinx")
	// 1. 创建server句柄
	s := znet.NewServer("[zinx v0.5]")
	// 2. 注册路由
	s.AddRouter(MsgPing, &PingRouter{})
	s.AddRouter(MsgHello, &HelloRouter{})

	// 3. 启动server
	s.Run()
}