package main

import (
	"TCPCache/server/cache"
	"TCPCache/server/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
