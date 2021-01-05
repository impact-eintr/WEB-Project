//负责master接受来自datanode的心跳信号
package heartbeat

import (
	"OSS_1.0/apiServer/rabbitmq"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var dataServers = make(map[string]time.Time)
var mutex sync.Mutex

//监视心跳信号
func ListenHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("apiServers")
	c := q.Consume()
	go removeExpiredDataServer()
	for msg := range c {
		dataServer, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

//扫描dataServers 并清除其中超过10s没收到心跳信号的数据节点
func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) { //已经有10秒没有更新
				delete(dataServers, s)
			}
		}
		mutex.Unlock()
	}
}

//遍历dataservers并返回所有当前所有的数据服务节点
func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	for s, _ := range dataServers {
		ds = append(ds, s)
	}
	log.Println("dataNode:", ds)
	return ds
}

//随机选择出一个节点
func ChooseRandomDataServer() string {
	ds := GetDataServers()
	n := len(ds)
	if n == 0 {
		return ""
	}
	dataNode := rand.Intn(n)
	return ds[dataNode]
}
