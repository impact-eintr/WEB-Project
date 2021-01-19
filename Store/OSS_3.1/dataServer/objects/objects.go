package objects

import (
	"OSS/dataServer/conf"
	"OSS/dataServer/locate"
	"OSS/dataServer/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Get(c *gin.Context) {
	file := getFile(strings.Split(c.Request.URL.EscapedPath(), "/")[2])
	if file == "" {
		c.Status(http.StatusNotFound)
		return
	}

	f, _ := os.Open(file)
	defer f.Close()
	data, _ := ioutil.ReadAll(f)

	c.Data(http.StatusOK, "application/octet-stream", data)

}

func getFile(hash string) string {
	file := conf.Conf.Dir + "/objects/" + hash
	f, _ := os.Open(file)
	d := url.PathEscape(utils.CalculateHash(f))
	f.Close()
	if d != hash {
		log.Println("object hash mismatch, remove", file)
		locate.Del(hash)
		os.Remove(file)
		return ""

	}
	return file

}
