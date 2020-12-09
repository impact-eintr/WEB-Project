package main

import (
	"TCPCache/server/cache"
	"TCPCache/server/http"
	"TCPCache/server/tcp"
)

func main() {
	ca := cache.New("inmemory")
	go tcp.New(ca).Listen()
	http.New(ca).Listen()
}
