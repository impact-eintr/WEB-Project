package cache

import (
	"log"
)

func New(typ string) Cache {
	var c Cache
	if typ == "inmemory" {
		c = newInMemoryCache()
	}

	if c == nil {
		panic("未知的缓存类型" + typ)
	}
	log.Println(typ, "read to serve")
	return c
}
