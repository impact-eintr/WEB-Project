package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/user", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.Run(":8081")
}
