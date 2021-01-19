package temp

import (
	"OSS/dataServer/conf"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Post(c *gin.Context) {
	output, _ := exec.Command("uuidgen").Output() //输出一个[]byte
	uuid := strings.TrimSuffix(string(output), "\n")
	//等价于uuid := string(output)[:len(output)-1]
	name := strings.Split(c.Request.URL.EscapedPath(), "/")[2]
	size, err := strconv.ParseInt(c.Request.Header.Get("size"), 0, 64)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	t := tempInfo{
		uuid,
		name,
		size,
	}

	err = t.writeToFile()
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	os.Create(conf.Conf.Dir + "/temp/" + t.Uuid + ".dat")
	c.Data(http.StatusOK, "application/octet-stream", []byte(uuid))
}

func (t *tempInfo) writeToFile() error {
	f, err := os.Create(conf.Conf.Dir + "/temp/" + t.Uuid)
	if err != nil {
		return err
	}
	defer f.Close()
	b, _ := json.Marshal(t)
	f.Write(b)
	return nil
}
