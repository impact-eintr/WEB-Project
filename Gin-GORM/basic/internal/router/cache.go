package router

import (
	"basic/internal/dao/webcache/cache"
	"basic/internal/middleware"
	"basic/pkg/cachehttp"

	"github.com/gin-gonic/gin"
)

func CacheRoute(r *gin.Engine, c cache.Cache) {
	cacheGroup := r.Group("/cache")
	{
		cacheGroup.Use(middleware.Cors(), middleware.PathParse)
		cacheGroup.Any("/hit/*key", cachehttp.New(c).CacheCheck, func(c *gin.Context) {
			miss := c.GetBool("miss") // 检查是否命中缓存
			if miss {
				c.Request.URL.Path = "/info" + c.Param("key") // 将请求的URL修改
				r.HandleContext(c)                            // 继续之后的操作

			}
		})

		cacheGroup.PUT("/update/*key", cachehttp.New(c).UpdateHandler)
		cacheGroup.GET("/status/", cachehttp.New(c).StatusHandler)

	}

}
