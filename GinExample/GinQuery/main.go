package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// /?key1=val1&key2=val2&key3=val3
	r.GET("/", func(c *gin.Context) {
		//name := c.Query("query")
		//name := c.DefaultQuery("query","none")
		name, ok := c.GetQuery("query")
		if !ok {
			log.Println("err")
			name = "none"
		}
		age, ok := c.GetQuery("age")
		if !ok {
			log.Println("err")
			name = "none"
		}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":8081")
}
