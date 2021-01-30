package cache

import "sync"

type inMemoryCache struct {
	c     map[string][]byte
	mutex sync.RWMutex
	Stat
}

func newInMemryCache() *inMemoryCache {
	return &inMemoryCache{
		make(map[string][]byte),
		sync.RWMutex{},
		Stat{},
	}
}

func (this *inMemoryCache) Set(k string, v []byte) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	temp, exist := this.c[k]
	if exist {
		this.del(k, temp) //更新
	}
	this.c[k] = v
	this.add(k, v)
	return nil
}

func (this *inMemoryCache) Del(k string) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	temp, exist := this.c[k]
	if exist {
		delete(this.c, k)
		this.del(k, temp) //更新
	}
	return nil
}

func (this *inMemoryCache) Get(k string) ([]byte, error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.c[k], nil

}

func (this *inMemoryCache) GetStat() Stat {
	return this.Stat
}
