package locate

import (
	"OSS/dataServer/conf"
	"OSS/dataServer/rabbitmq"
	"github.com/fatih/color"
	"path/filepath"
	"strconv"
	"sync"
)

var mutex sync.Mutex
var objects = make(map[string]int)

func Locate(hash string) bool {
	mutex.Lock()
	_, ok := objects[hash]
	mutex.Unlock()
	return ok
}

func Add(hash string) {
	mutex.Lock()
	objects[hash] = 1
	mutex.Unlock()
}

func Del(hash string) {
	mutex.Lock()
	delete(objects, hash)
	mutex.Unlock()
}

func StartLocate(url string) {
	q := rabbitmq.New(conf.Conf.RabbitmqAddr)
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		hash, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)

		}
		exist := Locate(hash)
		color.Yellow("服务节点请求数据:%v\t状态: %v\n", hash, exist)
		if exist {
			q.Send(msg.ReplyTo, url)
		}
	}
	color.Green("%v\n", objects)
}

func CollectObjects() {
	files, _ := filepath.Glob(conf.Conf.Dir + "/objects/*")
	for i := range files {
		hash := filepath.Base(files[i])
		objects[hash] = 1
	}
}
