package main

import (
	"OSS/dataServer/conf"
	"OSS/dataServer/heartbeat"
	"OSS/dataServer/locate"
	"OSS/dataServer/objects"
	"github.com/gin-gonic/gin"
)

func init() {
	confile := "conf/conf.json"
	conf.Conf.Parse(confile)
}

func main() {
	engine := gin.Default() //返回默认引擎

	var url string //监听地址:端口
	url = conf.Conf.ListenAddr + ":" + conf.Conf.ListenPort

	go heartbeat.StartHeartbeat(url)
	go locate.StartLocate(url)

	engine.GET("/objects", objects.Get)
	engine.PUT("/objects", objects.Put)

	engine.Run(url)
}
