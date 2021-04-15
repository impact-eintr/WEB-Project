package v1

import (
	"basic/internal/dao/webcache/cache"
)

type Server struct {
	cache.Cache
}

func NewServer(c cache.Cache) *Server {
	return &Server{c}
}
