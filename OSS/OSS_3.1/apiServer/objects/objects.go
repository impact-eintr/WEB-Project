package objects

import (
	"OSS/apiServer/es"
	"OSS/apiServer/heartbeat"
	"OSS/apiServer/locate"
	"OSS/apiServer/objectstream"
	"OSS/apiServer/utils"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"io"
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
	var e error
	if len(versionId) != 0 {
		version, e = strconv.Atoi(versionId)
		if e != nil {
			log.Println(e)
			c.Status(http.StatusBadRequest)
			return
		}
	}
	meta, e := es.GetMetadata(name, version)
	if e != nil {
		log.Println(e)
		c.Status(http.StatusInternalServerError)
		return

	}
	if meta.Hash == "" {
		c.Status(http.StatusNotFound)
		return

	}
	object := url.PathEscape(meta.Hash)
	stream, e := getStream(object)
	if e != nil {
		log.Println(e)
		c.Status(http.StatusNotFound)
		return

	}
	data, _ := ioutil.ReadAll(stream)
	c.Data(http.StatusOK, "application/octet-stream", data)
}

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)

	}
	return objectstream.NewGetStream(server, object)

}

func Put(c *gin.Context) {
	hash := utils.GetHashFromHeader(c.Request.Header) //从header中获取hash信息
	if hash == "" {
		log.Println("missing object hash in digest header")
		c.Status(http.StatusBadRequest)
		return
	}

	size := utils.GetSizeFromHeader(c.Request.Header)
	log.Println("来自客户端的PUT信息:")
	color.Yellow("RawHash : %v \nHash : %v \nSize : %v\n", hash, url.PathEscape(hash), strconv.FormatInt(size, 10))
	statuscode, err := storeObject(c.Request.Body, hash, size) //向dataserver发送带有hash and size信息的request
	if err != nil {
		log.Println(err)
		c.Status(statuscode)
		return
	}

	if statuscode != http.StatusOK {
		c.Status(statuscode)
		return
	}

	name := c.Param("file")
	err = es.AddVersion(name, hash, size)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
	}

}

func storeObject(r io.Reader, hash string, size int64) (int, error) {
	if locate.Exist(url.PathEscape(hash)) { //dataserver那里存储的是以hash值为文件名的文件
		return http.StatusOK, nil
	}

	stream, err := putStream(url.PathEscape(hash), size)
	if err != nil {
		return http.StatusServiceUnavailable, err
	}

	reader := io.TeeReader(r, stream)
	d := utils.CalculateHash(reader) //这里同时也调用了stream的Write()方法，向dataserver递送了PATCH请求

	color.Red("apiServer计算的这个是什么? : %v\n", d)
	if d != hash {
		stream.Commit(false)
		return http.StatusBadRequest, fmt.Errorf("object hash MISSMATCH , 文件计算hash : %s . 客户端提供的 : %s", d, hash)
	}

	stream.Commit(true)
	return http.StatusOK, nil
}

func putStream(object string, size int64) (*objectstream.TempPutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return objectstream.NewTempPutStream(server, object, size)
}

func Delete(c *gin.Context) {
	name := c.Param("file")
	version, e := es.SearchLatestVersion(name)
	if e != nil {
		log.Println(e)
		c.Status(http.StatusInternalServerError)
		return
	}

	e = es.PutMetadata(name, version.Version+1, 0, "")
	if e != nil {
		log.Println(e)
		c.Status(http.StatusInternalServerError)
		return
	}
}
