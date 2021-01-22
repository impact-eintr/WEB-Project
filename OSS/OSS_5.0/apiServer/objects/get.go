package objects

import (
	"OSS/apiServer/es"
	"OSS/apiServer/heartbeat"
	"OSS/apiServer/locate"
	"OSS/apiServer/rs"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func Get(c *gin.Context) {
	name := c.Param("file")
	versionId := c.Query("version")
	version := 0
	var err error

	if len(versionId) != 0 {
		version, err = strconv.Atoi(versionId)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusBadRequest)
			return
		}
	}

	meta, err := es.GetMetadata(name, version) //元数据服务
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	if meta.Hash == "" {
		c.Status(http.StatusNotFound)
		return
	}

	hash := url.PathEscape(meta.Hash)
	stream, err := GetStream(hash, meta.Size)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusNotFound)
		return
	}

	data, err := ioutil.ReadAll(stream)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusNotFound)
		return
	}
	c.Data(http.StatusOK, "application/octet-stream", data)
	stream.Close()
}

func GetStream(hash string, size int64) (*rs.RSGetStream, error) {
	LocateInfo := locate.Locate(hash)
	if len(LocateInfo) < rs.DATA_SHARDS {
		return nil, fmt.Errorf("object %s locate fail,result %v", hash, LocateInfo)

	}

	dataServers := make([]string, 0)
	if len(LocateInfo) != rs.ALL_SHARDS {
		dataServers = heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS-len(LocateInfo), LocateInfo)
	} //拿到足够的数据节点
	return rs.NewRSGetStream(LocateInfo, dataServers, hash, size)

}
