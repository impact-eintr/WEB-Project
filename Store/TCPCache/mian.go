package main

import (
	"github.com/Impact-yxw/WEB-Project/TCPCache/cache"
	"github.com/Impact-yxw/WEB-Project/TCPCache/http"
)

func main() {
	c := cache.New("inmemory")
	http.NewServer(c).Listen()
}
