package router

import (
	"basic/internal/middleware"
	"basic/pkg/cachehttp"

	"github.com/gin-gonic/gin"
)

func CacheRoute(r *gin.Engine, s *cachehttp.Server) {
	cacheGroup := r.Group("/cache")
	{
		cacheGroup.Use(middleware.Cors(), middleware.PathParse)
		cacheGroup.GET("/hit/*key", s.CacheCheck, func(c *gin.Context) {
			miss := c.GetBool("miss") // 检查是否命中缓存
			if miss {
				c.Request.URL.Path = "/info" + c.Param("key") // 将请求的URL修改
				r.HandleContext(c)                            // 继续之后的操作

			}
		})

		cacheGroup.PUT("/update/*key", s.UpdateHandler)
		cacheGroup.DELETE("/delete/*key", s.DeleteHandler)
		cacheGroup.GET("/status/", s.StatusHandler)

	}

}
