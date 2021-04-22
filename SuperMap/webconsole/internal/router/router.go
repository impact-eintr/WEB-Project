package router

import (
	"log"
	_ "webconsole/docs"
	"webconsole/global"
	"webconsole/internal/dao/webcache/cache"
	"webconsole/internal/middleware"
	"webconsole/pkg/logger"
	"webconsole/pkg/tcp"

	v1 "webconsole/internal/router/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() (r *gin.Engine, err error) {

	if global.ServerSetting.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r = gin.New()

	r.Use(logger.GinLogger())
	r.Use(logger.GinRecovery(true))
	r.Use(middleware.Cors())

	// 配置缓存服务
	c := cache.New(global.CacheSetting.CacheType, global.CacheSetting.TTL)
	// 开启缓存服务
	go tcp.New(c).Listen()

	s := v1.NewServer(c)
	info := v1.NewInfo()

	// 注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		// 缓存路由
		cacheGroup := apiv1.Group("/cache")
		{
			// 操作缓存
			cacheGroup.Use(middleware.PathParse)
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
			infoGroup.Use(middleware.PathParse)

			infoGroup.GET("/:infotype/:count",
				middleware.QueryRouter,
				info.GetUpdateInfo,
				func(c *gin.Context) {
					if c.GetString("type") == "mem" {
						log.Println("后续执行了")
						r.HandleContext(c) //继续之后的操作
					}
				})
		}
	}

	return r, nil
}
