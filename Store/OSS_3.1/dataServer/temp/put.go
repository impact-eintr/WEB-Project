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

func Put(c *gin.Context) {
	uuid := c.Param("tempfile")
	tempinfo, err := readFromFile(uuid)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusNotFound)
		return
	}

	infoFile := conf.Conf.Dir + "/temp/" + uuid
	dataFile := infoFile + ".dat"

	f, err := os.Open(dataFile)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	actual := info.Size()
	color.Yellow("临时文件请求的大小：%v\n", tempinfo.Size)
	color.Yellow("临时文件实际的大小：%v\n", actual)
	os.Remove(infoFile) //删除临时信息文件
	if actual != tempinfo.Size {
		os.Remove(dataFile)
		log.Println("actual size mismatch,expect ", tempinfo.Size, "actual", actual)
		c.Status(http.StatusInternalServerError)
		return
	}

}
