package cache

import (
	"log"
)

func New(typ string, ttl int) Cache {
	var c Cache

	if typ == "mem" {
		c = newInmemoryCache(ttl)
	} else if typ == "disk" {
		c = newRocksdbCache(ttl)
	}

	if c == nil {
		panic("未指定类型")
	}

	log.Println(typ, "服务已就位")
	return c
}
