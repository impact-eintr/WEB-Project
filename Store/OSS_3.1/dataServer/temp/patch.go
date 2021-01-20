package temp

import (
	"OSS/dataServer/conf"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func Patch(c *gin.Context) {
	uuid := c.Param("uuid")
	tempinfo, err := readFromFile(uuid)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusNotFound)
		return
	}

	infoFile := conf.Conf.Dir + "/temp/" + uuid
	dataFile := infoFile + ".dat"

	f, err := os.OpenFile(dataFile, os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, c.Request.Body) //向临时文件写入内容
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	info, err := f.Stat()
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	actual := info.Size()
	color.Yellow("临时文件请求的大小：%v\n", tempinfo.Size)
	color.Yellow("临时文件实际的大小：%v\n", actual)
	if actual > tempinfo.Size {
		os.Remove(infoFile)
		os.Remove(dataFile)
		log.Println("actual size", actual, "exceeds", tempinfo.Size)
		c.Status(http.StatusInternalServerError)
	}

}
