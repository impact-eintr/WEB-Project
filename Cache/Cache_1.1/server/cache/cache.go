package cache

import "log"

//相关的接口
type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStat() Stat
}

//相关的结构体
type Stat struct {
	Count     int64 `json:"count"`
	KeySize   int64 `json:"keysize"`
	ValueSize int64 `json:"valuesize"`
}

func New(typ string) Cache {
	var c Cache
	if typ == "inmemory" {
		c = newInMemryCache()
	}
	if c == nil {
		log.Fatalln("未知的类型")
	}
	log.Println(typ, "is ready to server")
	return c
}
