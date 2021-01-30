package http

import (
	"Cache/server/cache"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type cacheHandler struct {
	cache.Cache
}

func (this *cacheHandler) PutVal(c *gin.Context) {
	key := c.Param("key")
	if len(key) == 0 {
		c.Status(http.StatusInternalServerError)
		return
	}

	b, _ := ioutil.ReadAll(c.Request.Body)
	if len(b) != 0 {
		err := this.Set(key, b)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
	}

}

func (this *cacheHandler) GetStatus(c *gin.Context) {
	data, err := json.Marshal(this.GetStat())
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Data(http.StatusOK, "application/octet-stream", data)

}

func (this *cacheHandler) GetVal(c *gin.Context) {
	key := c.Param("key")
	b, err := this.Get(key)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Data(http.StatusOK, "application/octet-stream", b)
}

func (this *cacheHandler) DelVal(c *gin.Context) {
	key := c.Param("key")
	err := this.Del(key)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)

}
