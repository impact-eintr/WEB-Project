package main

import (
	"Cache/server/cache"
	"Cache/server/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("vim-go")
	engine := gin.Default()

	c := cache.New("inmemory")
	handler := http.New(c)

	engine.PUT("/cache/:key", handler.PutVal)
	engine.GET("/cache/:key", handler.GetVal)
	engine.DELETE("/cache/:key", handler.DelVal)

	engine.GET("/status", handler.GetStatus)

	engine.Run(":6077")
}
