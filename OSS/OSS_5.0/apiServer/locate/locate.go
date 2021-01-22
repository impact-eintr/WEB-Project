package locate

import (
	"OSS/apiServer/conf"
	"OSS/apiServer/rabbitmq"
	"OSS/apiServer/rs"
	"OSS/apiServer/types"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Get(c *gin.Context) {
	info := Locate(c.Param("file"))
	if len(info) == 0 {
		c.Status(http.StatusNotFound)
		return

	}
	b, _ := json.Marshal(info)
	res, _ := strconv.Unquote(string(b))
	c.JSON(http.StatusOK, res)
}

func Locate(name string) (LocateInfo map[int]string) {
	q := rabbitmq.New(conf.Conf.RabbitmqAddr)
	q.Publish("dataServers", name)
	c := q.Consume()

	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()

	LocateInfo = make(map[int]string)
	for i := 0; i < rs.ALL_SHARDS; i++ {
		msg := <-c //每条信息包含了拥有某个分片数据节点的地址和分片的id{addr:ip,id:int}
		if len(msg.Body) == 0 {
			return
		}

		var info types.LocateMessage
		json.Unmarshal(msg.Body, &info)
		LocateInfo[info.Id] = info.Addr

	}
	return

}

func Exist(name string) bool {
	return len(Locate(name)) >= rs.DATA_SHARDS //判断定位服务反馈消息数量超过4才支持定位
}
