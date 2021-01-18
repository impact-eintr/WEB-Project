package heartbeat

import (
	"OSS/dataServer/conf"
	"OSS/dataServer/rabbitmq"
	"time"
)

func StartHeartbeat(url string) {
	q := rabbitmq.New(conf.Conf.RabbitmqAddr)
	defer q.Close()
	for {
		q.Publish("apiServers", url) //给apiServer消息队列发消息
		time.Sleep(5 * time.Second)  //5秒后再发
	}
}
