package main

import "github.com/gin-gonic/gin"
import "io/ioutil"
import "log"
import "net/http"
import "os"
import "strconv"
import "strings"

func main() {
	r := gin.Default()

	//http://127.0.0.1:8080/ping
	//路径传参
	r.GET("/OSS/objects/*file", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"path":    strconv.Quote(c.Param("file")[1:]),
			"rawpath": strings.Split(c.Request.URL.EscapedPath(), "/")[3],
		})
	})

	//返回二进制文件
	r.GET("/picture/test.jpg", func(c *gin.Context) {
		img := c.Param("test.jpg")
		f, err := os.Open("./test.jpg")
		if err != nil {
			log.Println(err)
		}
		data, _ := ioutil.ReadAll(f)

		c.Data(http.StatusOK, img, data)
	})

	//POST 表单传值法
	r.POST("/POST", func(c *gin.Context) {
		title := c.PostForm("title")
		c.JSON(200, gin.H{
			"message": title,
		})
	})

	type Login struct {
		User     string `form:"user" xml:"user" json:"user" binding:"required"`
		Password string `form:"password" xml:"password" json:"password" binding:"required"`
	}
	//binding
	r.POST("/Login", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		log.Println(json)
		//c.JSON(http.StatusOK, gin.H{
		//	"User":     json.User,
		//	"Password": json.Password,
		//})
	})

	//Query 参数传值
	//http://localhost:8080/welcome?firstname=jiang&&lastname=kun
	r.GET("/welcome", func(c *gin.Context) {
		school := c.DefaultQuery("school", "TYUT")
		firstname := c.Query("firstname")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello!\n %s %s %s", school, firstname, lastname)
	})

	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}

//wait timeout
//http2
//keep alive 复用连接
