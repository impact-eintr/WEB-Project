package main

import (
	"Zinx/ziface"
	"Zinx/znet"
	"fmt"
)

type HeartBeatRouter struct {
	znet.BaseRouter
}

type CheckHeartBeatRouter struct {
	znet.BaseRouter
}

const (
	MsgHeartBeat      uint32 = 0
	MsgCheckHeartBeat uint32 = 1
)

// 检测心跳信号
func (p *HeartBeatRouter) PreHandle(request ziface.IRequest) {

	fmt.Println("recv from client:msgID = ", request.GetMsgID(),
		",data = ", string(request.GetMsgData()))
}
func (p *HeartBeatRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	fmt.Println("recv from client:msgID = ", request.GetMsgID(),
		",data = ", string(request.GetMsgData()))

	// 先读取客户端数据 再写回数据到客户端
	err := request.GetConnection().SendMsg(222, []byte("pong...pong..pong.\n"))
	if err != nil {
		fmt.Println("call back error")
	}
}

func (h *CheckHeartBeatRouter) Handle(request ziface.IRequest) {
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
	s := znet.NewServer("[zinx v0.8]")
	// 2. 注册路由
	// 接收心跳
	s.AddRouter(MsgHeartBeat, &HeartBeatRouter{})
	// 转发心跳
	s.AddRouter(MsgCheckHeartBeat, &CheckHeartBeatRouter{})

	// 3. 启动server
	s.Run()
}
