package znet

import (
	"Zinx/ziface"
)

type Request struct {
	// 已经和客户端建立好的连接
	conn ziface.IConnection

	// 客户端请求的数据
	data []byte
}

// 获取当前连接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// 获取请求的消息数据
func (r *Request) GetData() []byte {
	return r.data
}
