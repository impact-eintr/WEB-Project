package tcp

import (
	"Cache/server/cache"
	"net"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	l, err := net.Listen("tcp", ":9428")
	if err != nil {
		panic(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go s.process(c) //开启goroutine服务新的tcp连接
	}
}

func New(c cache.Cache) *Server {
	return &Server{c}
}
