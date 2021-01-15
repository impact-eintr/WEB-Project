package locate

import (
	"OSS_1.1/apiServer/rabbitmq"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

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

func Locate(name string) string {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
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
