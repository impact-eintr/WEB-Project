package temp

import (
	"OSS/dataServer/conf"
	"OSS/dataServer/locate"
	"OSS/dataServer/utils"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Put(c *gin.Context) {
	uuid := c.Param("uuid")
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

	commitTempObject(dataFile, tempinfo)
}

func (t *tempInfo) hash() string {
	s := strings.Split(t.Name, "_")
	return s[0]
}

func (t *tempInfo) id() int {
	s := strings.Split(t.Name, "_")
	id, _ := strconv.Atoi(s[1])
	return id
}

func commitTempObject(dataFile string, tempinfo *tempInfo) {
	f, err := os.Open(dataFile)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	d := url.PathEscape(utils.CalculateHash(f))
	color.Red("d是?:%v ", d)
	err = os.Rename(dataFile, conf.Conf.Dir+"/objects/"+tempinfo.Name+"_"+d)
	if err != nil {
		log.Println(err)
	}
	locate.Add(tempinfo.hash(), tempinfo.id())
}
