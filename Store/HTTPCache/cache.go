package main

import (
	"log"
	"sync"
)

//the Status od HTTPCache
type Stat struct {
	Count     int64
	KeySize   int64
	ValueSize int64
}

func (s *Stat) add(k string, v []byte) {
	s.Count++
	s.KeySize += int64(len(k))
	s.ValueSize += int64(len(v))
}

func (s *Stat) del(k string, v []byte) {
	s.Count--
	s.KeySize -= int64(len(k))
	s.ValueSize -= int64(len(v))
}

type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStat() Stat //return the status of HTTPCache
}

func New(typ string) Cache {
	var c Cache
	if typ == "inmemory" {
		c = newInMemoryCache()
	}
	log.Println(typ, "ready to serve")
	return c
}

type inMemoryCache struct {
	c     map[string][]byte
	mutex sync.RWMutex
	Stat  //继承Stat
}

func (c *inMemoryCache) Set(k string, v []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	tmp, exist := c.c[k]
	if exist {
		c.del(k, tmp)
	}
	c.c[k] = v
	c.add(k, v)
	return nil
}

func (c *inMemoryCache) Get(k string) ([]byte, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.c[k], nil

}

func (c *inMemoryCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	v, exist := c.c[k]
	if exist {
		delete(c.c, k) //delete is used to delete the element of map
		c.del(k, v)
	}
	return nil
}

func (c *inMemoryCache) GetStat() Stat {
	return c.Stat
}

//Create a new InMemoryCache
func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{
		make(map[string][]byte),
		sync.RWMutex{},
		Stat{},
	}
}
