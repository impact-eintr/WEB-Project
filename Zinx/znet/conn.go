package znet

import (
	"Zinx/ziface"
	"errors"
	"fmt"
	"io"
	"net"
)

type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn

	// 连接的ID
	ConnID uint32

	// 当前的连接状态
	isClosed bool

	// 该连接处理的方法Router
	Router ziface.IRouter

	// 告知当前连接已经退出/停止的channel
	ExitChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

// 连接读取数据的业务
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID:", c.ConnID, "Reader is exit, remote addr is ", c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		//buf := make([]byte, utils.GlobalConf.MaxPackageSize)
		//cnt, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("recv buf err", err)
		//	continue
		//}
		dp := NewDataPack()

		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error", err)
			break
		}
		// 解包 获取数据长度
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("read msg head error", err)
			break
		}

		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg head error", err)
				break
			}
		}

		msg.SetMsgData(data)

		// 得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			msg:  msg,
		}

		// 执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}
}

// 启动连接 让当前的连接准备开始工作
func (c *Connection) Start() {
	fmt.Println("Conn Start ()... ConnID:", c.ConnID)
	go c.StartReader()
}

// 停止连接 结束当前连接的动作
func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnID", c.ConnID)
	if c.isClosed {
		return
	}

	c.isClosed = true

	c.Conn.Close()

	close(c.ExitChan)
}

// 获取当前连接模块绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前连接模块的连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的TCP状态 IP port
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据 将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {
	return nil
}

// 提供一个SendMsg方法 将我们要发送到客户端的数据 先进行封包再发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when send msg")
	}

	dp := NewDataPack()
	msgpkg := NewMsgPackage(msgId, data)

	msg, err := dp.Pack(msgpkg)
	if err != nil {
		fmt.Println("Pack error msgid:", msgId)
		return errors.New("Pack error")
	}
	// 将数据发送给客户端
	if _, err = c.Conn.Write(msg); err != nil {
		fmt.Println("write msg error id", msgId, err)
	}

	return nil
}
