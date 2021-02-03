package process

import (
	"TCPDemo/client/common"
	"TCPDemo/client/module"
	"github.com/fatih/color"
)

//客户端维护的用户列表
var onlineUsers map[string]*common.User = make(map[string]*common.User, 1024)
var CurUser module.CurUser

//处理返回的NotifyUserStatusMes
func updateUserStatus(notifyStatusMes *common.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyStatusMes.Uid]
	if !ok {
		user = &common.User{
			Uid: notifyStatusMes.Uid,
		}
	}
	user.Ustatus = notifyStatusMes.UStatus

	onlineUsers[notifyStatusMes.Uid] = user
	//显示在线用户
	showOnlineUser()
}

func showOnlineUser() {
	for id := range onlineUsers {
		color.Blue("用户id:%s", id)
	}
}
