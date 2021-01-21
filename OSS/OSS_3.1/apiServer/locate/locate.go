package locate

import (
	"OSS/apiServer/conf"
	"OSS/apiServer/rabbitmq"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func Get(c *gin.Context) {
	info := Locate(url.PathEscape(c.Param("hash")[1:]))
	if len(info) == 0 {
		c.Status(http.StatusNotFound)
		return

	}
	b, _ := json.Marshal(info)
	url, _ := strconv.Unquote(string(b))
	c.JSON(http.StatusOK, url)
}

func Locate(name string) string {
	q := rabbitmq.New(conf.Conf.RabbitmqAddr)
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()

	}()
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s

}

func Exist(name string) bool {
	return Locate(name) != ""

}
