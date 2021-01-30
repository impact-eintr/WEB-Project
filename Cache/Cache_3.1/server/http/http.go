package http

import (
	"Cache/server/cache"
)

func New(c cache.Cache) *cacheHandler {
	return &cacheHandler{c}
}
