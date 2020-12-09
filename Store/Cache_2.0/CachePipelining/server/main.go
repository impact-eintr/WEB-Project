package main

import (
	"CachePipelining/server/cache"
	"CachePipelining/server/http"
	"CachePipelining/server/tcp"
	"flag"
	"log"
)

func main() {
	typ := flag.String("type", "inmemory", "cache type")
	flag.Parse()
	log.Println("type is ", *typ)
	ca := cache.New(*typ)
	go tcp.New(ca).Listen()
	http.New(ca).Listen()
}
