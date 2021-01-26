package temp

import (
	"OSS/dataServer/conf"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func Get(c *gin.Context) {
	uuid := url.PathEscape(c.Param("uuid"))
	f, err := os.Open(conf.Conf.Dir + "/temp/" + uuid + ".dat")
	if err != nil {
		log.Println(err)
		c.Status(http.StatusNotFound)
		return
	}
	defer f.Close()

	data, _ := ioutil.ReadAll(f)
	c.Data(http.StatusOK, "application/octet-stream", data)

}
