package objects

import (
	"OSS/dataServer/conf"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Put(c *gin.Context) {
	f, err := os.Create(conf.Conf.Dir + "/objects/" + c.Param("file"))
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, c.Request.Body)
}

func Get(c *gin.Context) {
	f, err := os.Open(conf.Conf.Dir + "/objects/" + c.Param("file"))
	log.Println("test")
	defer f.Close()

	if err != nil {
		log.Println(err)
		c.Status(http.StatusNotFound)
		return
	}
	data, _ := ioutil.ReadAll(f)

	c.Data(http.StatusOK, "application/octet-stream", data)
}
