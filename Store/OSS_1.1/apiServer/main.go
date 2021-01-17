package main

import (
	"OSS/apiServer/conf"
	"OSS/apiServer/heartbeat"
	"OSS/apiServer/locate"
	"OSS/apiServer/objects"
	"github.com/gin-gonic/gin"
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

	engine := gin.Default()

	//启动心跳 返回dataservers中随机一个地址
	go heartbeat.ListenHeartbeat()

	engine.PUT("/OSS/objects", objects.Put)
	engine.GET("/OSS/objects", objects.Get)
	engine.GET("/OSS/locate", locate.Get)
	engine.Run(url)
}
