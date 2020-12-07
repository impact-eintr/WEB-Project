package http

import (
	"github.com/impact-eintr/WEB-Project/Store_v2.0/HttpCache/cache"
	"net/http"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.cacheHandler())
	http.ListenAndServe(":12345", nil)
}

func New(c cache.Cache) *Server {
	return &Server{c}
}
