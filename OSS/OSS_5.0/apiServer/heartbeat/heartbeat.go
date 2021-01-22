package heartbeat

import (
	"OSS/apiServer/conf"
	"OSS/apiServer/rabbitmq"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var dataServers = make(map[string]time.Time)
var mutex sync.Mutex

func ListenHeartbeat() {
	q := rabbitmq.New(conf.Conf.RabbitmqAddr)
	defer q.Close()
	q.Bind("apiServers")
	c := q.Consume()
	go removeExpiredDataServer()
	for msg := range c {
		dataServer, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)

		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()

	}

}

func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mutex.Unlock()
	}

}

func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	for s, _ := range dataServers {
		ds = append(ds, s)
	}
	return ds

}

func ChooseRandomDataServers(n int, exclude map[int]string) (ds []string) {
	candidates := make([]string, 0)
	reverseExcludeMap := make(map[string]int)
	for id, Addr := range exclude {
		reverseExcludeMap[Addr] = id //反转键值对
	}

	servers := GetDataServers()

	for i := range servers {
		s := servers[i]
		_, excluded := reverseExcludeMap[s]
		if !excluded { //拥有一部分数据
			candidates = append(candidates, s)
		}
	}
	length := len(candidates)
	if length < n {
		return //返回空切片
	}

	p := rand.Perm(length)
	for i := 0; i < n; i++ {
		ds = append(ds, candidates[p[i]]) //至少4个拥有分片的服务器队列
	}

	return

}
