package znet

import (
	"Zinx/ziface"
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
	fmt.Println("成功启动server", s.Name)

	// 阻塞地等待客户端连接处理客户端连接业务
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}

		// 已经建立了连接 实现一个简单的回显业务
		go func() {
			for {

				buf := make([]byte, 512)
				cnt, err := conn.Read(buf)
				if err != nil {
					fmt.Println("recv buf err", err)
					continue
				}

				fmt.Println("recv buf:", string(buf))
				if _, err := conn.Write(buf[:cnt]); err != nil {
					fmt.Println("write buf err", err)
					continue
				}
			}
		}()

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
