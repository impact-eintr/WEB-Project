package main

import (
	"RocksDB/server/cache"
	"RocksDB/server/http"
	"RocksDB/server/tcp"
)

func main() {
	ca := cache.New("rocksdb")
	go tcp.New(ca).Listen()
	http.New(ca).Listen()
}
