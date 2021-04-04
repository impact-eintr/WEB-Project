package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./login.html", "./index.html")

	r.GET("/login", func(c *gin.Context) {
		//username := c.PostForm("username")
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name":     username,
			"password": password,
		})
	})

	r.Run(":8081")
}
