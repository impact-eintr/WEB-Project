package main

import (
	"basic/global"
	"basic/internal/dao/webcache/cache"
	"basic/internal/router"
	"basic/pkg/tcp"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	global.ConfigInit()

	// 配置缓存服务
	c := cache.New(global.G.CacheConfig.CacheType, global.G.CacheConfig.TTL)

	// 开启缓存服务
	go tcp.New(c).Listen()

	r := gin.Default()

	// 缓存路由组
	router.CacheRoute(r, c)
	// 数据查询路由组
	router.InfoRoute(r)

	r.Run("0.0.0.0:8081")
}
