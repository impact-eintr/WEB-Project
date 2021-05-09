package znet

import (
	"Zinx/ziface"
	"fmt"
	"log"
)

// 消息管理实现
type MsgHandle struct {
	// 存放每个MsgID所对应的处理方法(之后是否可以编写路由组)
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

// 调度/执行对应的Router消息处理方法
func (h *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	// 1. 从Request中找到msgID
	handler, ok := h.Apis[request.GetMsgID()]
	if !ok {
		fmt.Printf("api msg[%v] is NOT FOUND! Need Register", request.GetMsgID())
	}

	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)

}

// 为消息添加具体的处理逻辑
func (h *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := h.Apis[msgId]; ok {
		log.Println("repeat api, migId = ", msgId)
		return
	}

	h.Apis[msgId] = router
	fmt.Println("Add Router [MsgId] Succ")

}
