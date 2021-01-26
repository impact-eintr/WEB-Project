package locate

import (
	"OSS/dataServer/conf"
	"OSS/dataServer/rabbitmq"
	"OSS/dataServer/types"
	"github.com/fatih/color"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var mutex sync.Mutex
var objects = make(map[string]int)

func Locate(hash string) int {
	mutex.Lock()
	id, ok := objects[hash]
	mutex.Unlock()
	if !ok {
		return -1
	}
	return id
}

func Add(hash string, id int) {
	mutex.Lock()
	objects[hash] = id
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

		id := Locate(hash)
		color.Yellow("服务节点请求数据:%v\t状态: %v\n", hash, id)
		if id != -1 {
			q.Send(msg.ReplyTo, types.LocateMessage{Addr: url, Id: id})

		}
	}
	color.Green("%v\n", objects)
}

func CollectObjects() {
	files, _ := filepath.Glob(conf.Conf.Dir + "/objects/*")
	for i := range files {
		file := strings.Split(filepath.Base(files[i]), "_")
		color.Yellow("文件切分出个什么？ :%v\n", file)

		if len(file) != 3 {
			panic(file[i])
		}

		hash := file[0]
		id, err := strconv.Atoi(file[1])
		if err != nil {
			panic(err)
		}
		objects[hash] = id
	}
}
