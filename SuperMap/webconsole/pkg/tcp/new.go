package tcp

import (
	"net"
	"webconsole/global"
	"webconsole/internal/dao/webcache/cache"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	l, err := net.Listen("tcp", ":"+global.CacheSetting.Port)
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
