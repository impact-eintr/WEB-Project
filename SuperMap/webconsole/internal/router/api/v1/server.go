package v1

import (
	"webconsole/internal/dao/webcache/cache"
)

type Server struct {
	cache.Cache
}

func NewServer(c cache.Cache) *Server {
	return &Server{c}
}
