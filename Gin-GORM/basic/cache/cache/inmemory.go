package cache

import (
	"sync"
	"time"
)

type value struct {
	c       []byte
	created time.Time
}

type inMemoryCache struct {
	c     map[string]value //缓存键值对
	mutex sync.RWMutex     //读写一致性控制
	Stat                   //缓存当前状态
	ttl   time.Duration    //缓存生存时间
}

func newInmemoryCache(ttl int) *inMemoryCache {
	c := &inMemoryCache{
		make(map[string]value),
		sync.RWMutex{},
		Stat{},
		time.Duration(ttl) * time.Second,
	}
	if ttl > 0 {
		// 开启一个groutine后台处理缓存TTl
		go c.expirer()
	}
	return c
}

func (c *inMemoryCache) expirer() {
	for {
		time.Sleep(c.ttl)
		c.mutex.RLock()

		for k, v := range c.c {
			c.mutex.RUnlock()
			if v.created.Add(c.ttl).Before(time.Now()) {
				c.Del(k)
			}
			c.mutex.RLock()
		}

		c.mutex.RUnlock()
	}
}

func (c *inMemoryCache) Set(k string, v []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	tmp, exist := c.c[k]
	if exist {
		c.del(k, tmp.c)
	}
	c.c[k] = value{v, time.Now()}
	c.add(k, v)
	return nil
}

func (c *inMemoryCache) Get(k string) ([]byte, error) {
	c.mutex.RLock()
	val := c.c[k].c
	c.c[k] = value{val, time.Now()}

	defer c.mutex.RUnlock()
	return val, nil
}

func (c *inMemoryCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	v, exist := c.c[k]
	if exist {
		delete(c.c, k)
		c.del(k, v.c)
	}
	return nil
}

func (c *inMemoryCache) GetStat() Stat {
	return c.Stat
}
