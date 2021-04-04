package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func func1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}

func main() {

	r := gin.Default()

	r.GET("/", func1)
	r.GET("/search/", func1)
	r.GET("/support/", func1)
	r.GET("/blog/", func1)
	r.GET("/blog/:post/", func1)
	r.GET("/about-us/", func1)
	r.GET("/about-us/team/", func1)

	r.Run(":8080")

}
