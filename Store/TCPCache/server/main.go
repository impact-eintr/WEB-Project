package main

import (
	"github.com/impact-eintr/WEB-Project/TCPCache/cache"
	"github.com/impact-eintr/WEB-Project/TCPCache/http"
)

func main() {
	c := cache.New("inmemory")
	http.NewServer(c).Listen()
}
