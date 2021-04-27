package znet

import (
	"Zinx/ziface"
	"fmt"
	"net"
)

type Connection struct {
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn

	// 连接的ID
	ConnID uint32

	// 当前的连接状态
	isClosed bool

	// 当前连接所绑定的处理业务方法API
	handleAPI ziface.HandleFunc

	// 告知当前连接已经退出/停止的channel
	ExitChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: callback_api,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

// 连接读取数据的业务
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID:", c.ConnID, "Reader is exit, remote addr is ", c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Printf("ConnID[%v]\n handle err:%v", c.ConnID, err)
			break
		}
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
	if c.isClosed == true {
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
