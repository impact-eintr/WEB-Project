package locate

import (
	"OSS/dataServer/conf"
	"OSS/dataServer/rabbitmq"
	"os"
	"strconv"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate(url string) {
	q := rabbitmq.New(conf.Conf.RabbitmqAddr)
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		object, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)

		}
		if Locate(conf.Conf.Dir + "/objects/" + object) {
			q.Send(msg.ReplyTo, url)
		}
	}
}
