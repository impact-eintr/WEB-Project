package main

import (
	"Cache/server/cache"
	"Cache/server/http"
	"Cache/server/tcp"
	"flag"
	"log"
)

func main() {

	typ := flag.String("type", "inmemory", "cache type")
	flag.Parse()
	log.Println("type is ", *typ)
	c := cache.New(*typ)
	go tcp.New(c).Listen()
	http.New(c).Listen()
}
