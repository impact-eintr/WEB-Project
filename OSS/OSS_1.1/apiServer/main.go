package main

import (
	"OSS/apiServer/conf"
	"OSS/apiServer/heartbeat"
	"OSS/apiServer/locate"
	"OSS/apiServer/objects"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	confile := "conf/conf.json"
	conf.Conf.Parse(confile)
}

func main() {

	var url string //监听地址:端口
	url = conf.Conf.ListenAddr + ":" + conf.Conf.ListenPort
	log.Println(url)
	engine := gin.Default()

	//启动心跳 返回dataservers中随机一个地址
	go heartbeat.ListenHeartbeat()

	engine.PUT("/OSS/objects/:file", objects.Put)
	engine.GET("/OSS/objects/:file", objects.Get)
	engine.GET("/OSS/locate/:file", locate.Get)
	engine.Run(url)
}
