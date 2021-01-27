package process

import (
	"fmt"
)

type UserList struct {
	onlineUsers map[string]*UserProcess
}

//用户列表全局变量
var userList *UserList

func init() {
	userList = &UserList{
		onlineUsers: make(map[string]*UserProcess, 1024),
	}
}

func (this *UserList) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.Uid] = up
}

func (this *UserList) DelOnlineUser(uid string) {
	delete(this.onlineUsers, uid)
}

func (this *UserList) GetAllOnlineUser() map[string]*UserProcess {
	return this.onlineUsers
}

func (this *UserList) GetAllOnlineUserById(uid string) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[uid]
	if !ok {
		err = fmt.Errorf("用户%s 不存在", uid)
		return
	}
	return
}
