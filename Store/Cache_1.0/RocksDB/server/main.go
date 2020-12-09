package main

import (
	"RocksDB/server/cache"
	"RocksDB/server/http"
	"RocksDB/server/tcp"
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
