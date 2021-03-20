package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	name, _ := c.Get("name") //从上下文中取值（跨中间件取值）
	c.JSON(http.StatusOK, gin.H{
		"msg":  "This is the index page",
		"name": name,
	})
}

func homeHandler(c *gin.Context) {
	name, _ := c.Get("name")
	c.JSON(http.StatusOK, gin.H{
		"msg":  "This is the home page",
		"name": name,
	})
}

func m1(c *gin.Context) {
	fmt.Println("m1 is working...")
	start := time.Now()
	c.Next() //调用后续的处理函数 即indexHandler...
	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 exited...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 is working...")
	c.Next() //调用后续的处理函数 即indexHandler...
	fmt.Println("m2 exited...")
}

func m3(c *gin.Context) {
	c.Set("name", "eintr") //在上下文中设置值
}

func main() {
	r := gin.Default() //默认使用了Logger() 和 Recover()中间件
	//r := gin.New() //不包含任何中间件
	//当中间件或者handler中启动新的goroutine时不能使用原始上下文(c *gin.Context)必须使用其只读副本(c.Copy())
	userGroup := r.Group("/user")

	//为路由组设置中间件
	userGroup.Use(m1, m2, m3)

	userGroup.GET("/index", indexHandler)
	userGroup.GET("/home", homeHandler)

	r.Run(":8081")
}
