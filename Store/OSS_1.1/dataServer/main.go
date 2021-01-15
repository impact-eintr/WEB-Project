package main

import (
	"OSS_1.1/dataServer/heartbeat"
	"OSS_1.1/dataServer/locate"
	"OSS_1.1/dataServer/objects"
	"flag"
	"log"
	"net/http"
)

func main() {
	var rabbitmqAddr string
	var listenAddr string //监听主机地址
	var listenPort string //监听端口
	var url string        //监听地址:端口
	var dir string        //工作地址
	flag.StringVar(&rabbitmqAddr, "q", "amqp://test:test@192.168.4.1:5672", "消息队列机地址，默认为ubuntu")
	flag.StringVar(&listenAddr, "h", "127.0.0.1", "主机地址，默认为本机")
	flag.StringVar(&listenPort, "p", "54321", "主机地址，默认为本机")
	flag.StringVar(&dir, "r", "/tmp/objects", "存储目录，默认/tmp/objects")

	flag.Parse()
	url = listenAddr + ":" + listenPort

	go heartbeat.StartHeartbeat(rabbitmqAddr, url)
	go locate.StartLocate(rabbitmqAddr, url, dir)
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(url, nil))
}
