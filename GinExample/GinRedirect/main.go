package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	//请求重定向
	r.GET("index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com")
	})

	//请求转发
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b" //将请求的URL修改
		r.HandleContext(c)        //继续之后的操作
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	r.Run(":8081")
}
