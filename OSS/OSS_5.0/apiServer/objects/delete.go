package objects

import (
	"OSS/apiServer/es"
	"log"
	"net/http"
)

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
