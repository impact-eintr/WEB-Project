package main

import (
	"github.com/server/cache"
	"github.com/server/http"
	"github.com/server/tcp"
)

func main() {
	ca := cache.New("inMemory")
	go tcp.New(ca).Listen()
	http.New(ca).Listen()
}
