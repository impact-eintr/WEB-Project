package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	type msg struct {
		Name string `json:"name"`
		Word string `json:"word"`
	}

	r.GET("/json", func(c *gin.Context) {
		data := msg{
			Name: "eintr",
			Word: "wdnm!",
		}

		c.JSON(http.StatusOK, data)

	})

	r.Run(":8080")

}
