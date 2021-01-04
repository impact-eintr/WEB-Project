package heartbeat

import (
	"OSS_1.0/dataServer/rabbitmq"
	"os"
	"time"
)

func StartHeartBeat() {
	q := rabbitmq.New(os.Getenv("RABBITMA_SERVER"))
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)

	}
}
