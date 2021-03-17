package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("user/search", func(c *gin.Context) {
		username := c.PostForm("username")

	})

}
