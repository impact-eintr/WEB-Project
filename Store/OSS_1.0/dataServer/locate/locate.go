package locate

import (
	"OSS_1.0/dataServer/rabbitmq"
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
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		object, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}
		if Locate(os.Getenv("STORE_ROOT") + "/objects/" + object) {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}
}
