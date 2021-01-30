package main

import (
	"Cache/server/cache"
	"Cache/server/http"
	"Cache/server/tcp"
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

var cacheType *string

func init() {
	cacheType = flag.String("t", "inmemory", "cache type")
	flag.Parse()
	log.Println("type is ", *cacheType)
}

func main() {
	engine := gin.Default()

	c := cache.New(*cacheType)

	go tcp.New(c).Listen()

	handler := http.New(c)
	engine.PUT("/cache/:key", handler.PutVal)
	engine.GET("/cache/:key", handler.GetVal)
	engine.DELETE("/cache/:key", handler.DelVal)

	engine.GET("/status", handler.GetStatus)

	engine.Run(":6077")
}
