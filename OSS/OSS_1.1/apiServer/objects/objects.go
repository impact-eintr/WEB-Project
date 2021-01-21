package objects

import (
	"OSS/apiServer/heartbeat"
	"OSS/apiServer/locate"
	"OSS/apiServer/objectstream"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func get(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[3]
	stream, e := getStream(object)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	io.Copy(w, stream)

}

func Get(c *gin.Context) {
	object := c.Param("file")
	log.Println(object)
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
	fmt.Println("test")
	object := c.Param("file")
	fmt.Println(object)
	statuscode, e := storeObject(c.Request.Body, object)
	if e != nil {
		log.Println(e)
	}
	c.Status(statuscode)
}

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return objectstream.NewPutStream(server, object), nil
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
