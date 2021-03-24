package main

import (
	"Cache/server/cache"
	"Cache/server/http"
	"Cache/server/tcp"
)

func main() {
	c := cache.New("inmemory")
	go tcp.New(c).Listen()
	http.New(c).Listen()
}
