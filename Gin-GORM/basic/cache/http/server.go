package http

import (
	"basic/cache/cache"
)

type Server struct {
	cache.Cache
}

func New(c cache.Cache) *Server {
	return &Server{c}
}
