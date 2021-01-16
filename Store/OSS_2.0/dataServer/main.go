package main

import (
	"OSS/dataServer/conf"
	"OSS/dataServer/heartbeat"
	"OSS/dataServer/locate"
	"OSS/dataServer/objects"
	"log"
	"net/http"
)

func init() {
	confile := "conf/conf.json"
	conf.Conf.Parse(confile)
}

func main() {
	var url string //监听地址:端口
	url = conf.Conf.ListenAddr + ":" + conf.Conf.ListenPort
	log.Println(url)

	go heartbeat.StartHeartbeat(conf.Conf.RabbitmqAddr, url)
	go locate.StartLocate(conf.Conf.RabbitmqAddr, url, conf.Conf.Dir)
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(url, nil))
}
