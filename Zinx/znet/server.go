package znet

import (
	"Zinx/utils"
	"Zinx/ziface"
	"fmt"
	"log"
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
	// 当前的Server添加一个router server注册的连接对应的处理业务
	Router ziface.IRouter
}

func (s *Server) Start() {
	log.Println("[Zinx] Conf:", utils.GlobalConf)
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

		connDeal := NewConnection(conn, cid, s.Router)

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

// 添加router模块
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("[Zinx] add Router Suss!!")
}

// 初始化Server模块
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      utils.GlobalConf.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalConf.Host,
		Port:      utils.GlobalConf.Port,
		Router:    nil,
	}

	return s
}
