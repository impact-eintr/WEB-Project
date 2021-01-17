package heartbeat

import (
	"OSS_1.1/dataServer/rabbitmq"
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
