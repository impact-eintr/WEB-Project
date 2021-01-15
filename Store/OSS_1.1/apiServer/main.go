package main

import (
	"OSS_1.1/apiServer/heartbeat"
	"OSS_1.1/apiServer/locate"
	"OSS_1.1/apiServer/objects"
	"flag"
	"log"
	"net/http"
)

func main() {

	var rabbitmqAddr string
	var listenAddr string //监听主机地址
	var listenPort string //监听端口
	var url string        //监听地址:端口
	flag.StringVar(&rabbitmqAddr, "q", "127.0.0.1", "消息队列机地址，默认为本机")
	flag.StringVar(&listenAddr, "h", "127.0.0.1", "主机地址，默认为本机")
	flag.StringVar(&listenPort, "p", "12345", "主机地址，默认为本机")

	flag.Parse()
	url = listenAddr + ":" + listenPort
	go heartbeat.ListenHeartbeat(rabbitmqAddr)
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatalln(http.ListenAndServe(url, nil))

}
