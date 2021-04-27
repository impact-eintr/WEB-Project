package znet

import (
	"Zinx/ziface"
	"errors"
	"fmt"
	"net"
)

type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的IP
	IP string
	// 服务器监听的端口
	Port int
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listening on IP: %s, Port :%d, is starting\n", s.IP, s.Port)
	// 获取一个TCP连接的Addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("处理tcp地址出错:", err)
		return
	}

	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen ", s.IPVersion, "err", err)
	}

	// 初始化connid
	var cid uint32 = 0

	fmt.Println("成功启动server", s.Name)

	// 阻塞地等待客户端连接处理客户端连接业务
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}

		connDeal := NewConnection(conn, cid, func(conn *net.TCPConn, data []byte, cnt int) error {
			fmt.Println("[ConnHandle] CallbackToClient")
			if _, err := conn.Write(data[:cnt]); err != nil {
				fmt.Println("write back buf err", err)
				return errors.New("CallBackToClient error")

			}
			return nil
		})

		go connDeal.Start()

		cid++
	}
}

func (s *Server) Stop() {
	// TODO
}

func (s *Server) Run() {
	s.Start()

	// TODO

	// 阻塞状态
	select {}

}

// 初始化Server模块
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}

	return s
}
