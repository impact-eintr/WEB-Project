package cachehttp

import (
	"basic/internal/dao/webcache/cache"
)

type Server struct {
	cache.Cache
}

func New(c cache.Cache) *Server {
	return &Server{c}
}
