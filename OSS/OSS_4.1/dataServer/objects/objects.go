package objects

import (
	"OSS/dataServer/conf"
	"OSS/dataServer/locate"
	"crypto/sha256"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func Get(c *gin.Context) {
	file := getFile(url.PathEscape(c.Param("file")[1:]))
	if file == "" {
		c.Status(http.StatusNotFound)
		return
	}

	f, _ := os.Open(file)
	defer f.Close()
	data, _ := ioutil.ReadAll(f)

	c.Data(http.StatusOK, "application/octet-stream", data)

}

func sendFile(w io.Writer, file string) {
	f, _ := os.Open(file)
	defer f.Close()
	io.Copy(w, f)
}

func getFile(name string) string {
	files, _ := filepath.Glob(conf.Conf.Dir + "/objects/" + name + "_*")
	if len(files) != 1 {
		return ""

	}
	file := files[0]
	h := sha256.New()
	sendFile(h, file)
	d := url.PathEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	hash := strings.Split(file, "_")[2]
	if d != hash {
		log.Println("object hash mismatch, remove", file)
		locate.Del(hash)
		os.Remove(file)
		return ""

	}
	return file

}
