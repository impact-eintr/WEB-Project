package main

import (
	"OSS/dataServer/conf"
	"OSS/dataServer/heartbeat"
	"OSS/dataServer/locate"
	"OSS/dataServer/objects"
	"OSS/dataServer/temp"
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

	locate.CollectObjects()

	go heartbeat.StartHeartbeat(url)
	go locate.StartLocate(url)

	engine.GET("/objects/*file", objects.Get)
	engine.PUT("/temp/:uuid", temp.Put)
	engine.PATCH("/temp/:uuid", temp.Patch)
	engine.POST("/temp/*tempfile", temp.Post)
	engine.DELETE("/temp/:uuid", temp.Delete)

	engine.Run(url)
}
