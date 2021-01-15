package main

import (
	"OSS_1.1/apiServer/conf"
	"OSS_1.1/apiServer/heartbeat"
	"OSS_1.1/apiServer/locate"
	"OSS_1.1/apiServer/objects"
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

	go heartbeat.ListenHeartbeat(conf.Conf.RabbitmqAddr)
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatalln(http.ListenAndServe(url, nil))

}
