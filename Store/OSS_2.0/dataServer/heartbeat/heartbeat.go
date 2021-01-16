package heartbeat

import (
	"OSS/dataServer/rabbitmq"
	"time"
)

func StartHeartbeat(rabbitmqAddr string, listenAddr string) {
	q := rabbitmq.New(rabbitmqAddr)
	defer q.Close()
	for {
		q.Publish("apiServers", listenAddr)
		time.Sleep(5 * time.Second)
	}
}
