package locate

import (
	"OSS/dataServer/conf"
	"OSS/dataServer/rabbitmq"
	"path/filepath"
	"strconv"
	"sync"
)

var objects = make(map[string]int)
var mutex sync.Mutex

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
		if exist {
			q.Send(msg.ReplyTo, conf.Conf.RabbitmqAddr)
		}
	}
}

func CollectObjects() {
	files, _ := filepath.Glob(conf.Conf.Dir + "/objects/*")
	for i := range files {
		hash := filepath.Base(files[i])
		objects[hash] = 1
	}
}