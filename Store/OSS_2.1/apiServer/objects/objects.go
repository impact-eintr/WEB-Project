package objects

import (
	"OSS/apiServer/es"
	"OSS/apiServer/heartbeat"
	"OSS/apiServer/locate"
	"OSS/apiServer/objectstream"
	"OSS/apiServer/utils"
	"fmt"
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
	hash := utils.GetHashFromHeader(c.Request.Header)
	if hash == "" {
		log.Println("missing object hash in digest header")
		c.Status(http.StatusBadRequest)
		return
	}

	statuscode, e := storeObject(c.Request.Body, url.PathEscape(hash))
	if e != nil {
		log.Println(e)
		c.Status(statuscode)
		return
	}

	if statuscode != http.StatusOK {
		c.Status(statuscode)
		return
	}
	name := c.Param("file")
	size := utils.GetSizeFromHeader(c.Request.Header)
	e = es.AddVersion(name, hash, size)
	if e != nil {
		log.Println(e)
		c.Status(http.StatusInternalServerError)
	}
}

func storeObject(r io.Reader, object string) (int, error) {
	stream, e := putStream(object)
	if e != nil {
		return http.StatusServiceUnavailable, e
	}
	io.Copy(stream, r)
	e = stream.Close()
	if e != nil {
		return http.StatusInternalServerError, e
	}

	return http.StatusOK, nil
}

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return objectstream.NewPutStream(server, object), nil
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
