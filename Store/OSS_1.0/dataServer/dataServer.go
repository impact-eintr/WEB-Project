package main

import (
	"OSS_1.0/dataServer/Objects"
	"OSS_1.0/dataServer/heartbeat"
	"OSS_1.0/dataServer/locate"
	"log"
	"net/http"
	"os"
)

func main() {

	go heartbeat.StartHeartBeat() //开启数据节点的心跳服务
	go locate.StartLocate()       //开启数据节点的定位服务
	http.HandleFunc("/Objects/", Objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))

}
