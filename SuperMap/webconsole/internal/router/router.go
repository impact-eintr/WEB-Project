package router

import (
	_ "webconsole/docs"
	"webconsole/global"
	"webconsole/internal/dao/webcache/cache"
	"webconsole/internal/middleware"
	"webconsole/pkg/tcp"

	v1 "webconsole/internal/router/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	// 调用gin自带的日志收集 之后可以替换
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 配置缓存服务
	c := cache.New(global.CacheSetting.CacheType, global.CacheSetting.TTL)
	s := v1.NewServer(c)
	// 开启缓存服务
	go tcp.New(c).Listen()

	info := v1.NewInfo()

	// 注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		// 缓存路由
		cacheGroup := apiv1.Group("/cache")
		{
			// 操作缓存
			cacheGroup.Use(middleware.Cors(), middleware.PathParse)
			cacheGroup.GET("/hit/*key", s.CacheCheck, func(c *gin.Context) {
				miss := c.GetBool("miss") // 检查是否命中缓存
				if miss {
					c.Request.URL.Path = "/api/v1/info" + c.Param("key") // 将请求的URL修改
					r.HandleContext(c)                                   // 继续之后的操作

				}
			})

			cacheGroup.PUT("/update/*key", s.UpdateHandler)
			cacheGroup.DELETE("/delete/*key", s.DeleteHandler)

			// 获取缓存状态
			cacheGroup.GET("/status/", s.StatusHandler)
		}

		// 数据查询路由
		infoGroup := apiv1.Group("/info")
		{
			infoGroup.Use(middleware.Cors(), middleware.PathParse)

			infoGroup.GET("/:infotype/:count",
				middleware.QueryRouter,
				info.GetUpdateInfo,
				func(c *gin.Context) {
					r.HandleContext(c) //继续之后的操作
				})
		}
	}

	return r
}
