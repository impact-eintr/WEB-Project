package locate

import (
	"OSS_1.0/dataServer/rabbitmq"
	"log"
	"os"
	"strconv"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate() {
	q := rabbitmq.New(os.Getenv("RABBIT_SERVER"))
	defer q.Close()
	q.Bind("dataServers") //与dataServer绑定
	c := q.Consume()      //消息通道
	for msg := range c {
		log.Println(string(msg.Body))
		object, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}
		if Locate(os.Getenv("STORE_ROOT") + "/objects/" + object) {
			//向消息的发送方返回本服务的监听地址，表示该对象讯在本服务器节点上
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}
}
