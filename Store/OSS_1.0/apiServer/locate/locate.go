package locate

import (
	"OSS_1.0/apiServer/rabbitmq"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//处理HTTP请求 不处理GET以外的请求方法
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	info := Locate(strings.Split(r.URL.EscapedPath(), "/")[2])
	if len(info) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(info)
	w.Write(b)
}

//name是需要定位的对象的名字 创建一个新的消息队列 并向dataServers exchange群发这个对象的定位信息
func Locate(name string) string {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	q.Publish("dataServers", name) //群发
	c := q.Consume()               //群发后接收返回信息
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	msg := <-c
	log.Println(string(msg.Body))
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

//判断是否存在
func Exist(name string) bool {
	return Locate(name) != ""
}
