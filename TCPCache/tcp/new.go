package tcp

import (
	"github.com/TCPCache/cache"
	"net"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		panic(e)
	}
	for {
		c, e := l.Accept()
		if e != nil {
			panic(e)
		}

	}
}

func New(c cache.Cache) *Server {
	return Server{c}
}
