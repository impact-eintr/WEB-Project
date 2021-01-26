package objects

import (
	"OSS/apiServer/es"
	"OSS/apiServer/heartbeat"
	"OSS/apiServer/locate"
	"OSS/apiServer/rs"
	"OSS/apiServer/utils"
	//"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

//大文件专用接口
func Post(c *gin.Context) {
	name := c.Param("file")
	size, e := strconv.ParseInt(c.Request.Header.Get("size"), 0, 64)
	if e != nil {
		log.Println(e)
		c.Status(http.StatusForbidden)
		return
	}

	hash := utils.GetHashFromHeader(c.Request.Header)
	if hash == "" {
		log.Println("missing object hash in digest header")
		c.Status(http.StatusBadRequest)
		return
	}

	if locate.Exist(url.PathEscape(hash)) {
		e = es.AddVersion(name, hash, size) //重复的话，更新版本(版本号+1)
		if e != nil {
			log.Println(e)
			c.Status(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusOK)
		}
		return
	}

	ds := heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS, nil)
	if len(ds) != rs.ALL_SHARDS {
		log.Println("cannot find enough dataServer")
		c.Status(http.StatusServiceUnavailable)
		return
	}

	stream, e := rs.NewRSResumablePutStream(ds, name, url.PathEscape(hash), size) //创建post请求给datanode
	if e != nil {
		log.Println(e)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Header("location", "/temp/"+url.PathEscape(stream.ToToken()))
	c.Status(http.StatusCreated)

}
