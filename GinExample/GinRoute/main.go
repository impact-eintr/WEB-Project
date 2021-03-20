package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{
				"method": "PUT",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method": "POST",
			})
		case http.MethodDelete:
			c.JSON(http.StatusOK, gin.H{
				"method": "Delete",
			})
		}
	})

	// 路由组
	// 将共用的前缀提取出来 创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "video index",
			})
		})
		videoGroup.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "video home",
			})
		})
	}

	// 非法路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "-_-# 页面找不到呀",
		})
	})
	r.Run(":8081")

}
