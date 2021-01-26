package temp

import (
	"OSS/dataServer/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func Head(c *gin.Context) {
	uuid := c.Param("uuid")
	f, e := os.Open(conf.Conf.Dir + "/temp/" + uuid + ".dat")
	if e != nil {
		log.Println(e)
		c.Status(http.StatusNotFound)
		return

	}
	defer f.Close()
	info, e := f.Stat()
	if e != nil {
		log.Println(e)
		c.Status(http.StatusInternalServerError)
		return

	}
	c.Header("content-length", fmt.Sprintf("%d", info.Size()))

}
