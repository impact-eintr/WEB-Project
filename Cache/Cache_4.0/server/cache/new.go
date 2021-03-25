package cache

import (
	"log"
)

func New(typ string) Cache {
	var c Cache

	if typ == "inmemory" {
		c = newInmemoryCache()
	} else if typ == "rocksdb" {
		c = newRocksdbCache()
	}

	if c == nil {
		panic("未指定类型")
	}

	log.Println(typ, "服务已就位")
	return c
}
