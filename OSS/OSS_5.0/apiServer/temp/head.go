package temp

import (
	"OSS/apiServer/rs"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
)

func Head(c *gin.Context) {
	token := url.PathEscape(c.Param("file")[1:])
	log.Println(token)
	stream, e := rs.NewRSResumablePutStreamFromToken(token)
	if e != nil {
		log.Println(e)
		c.Status(http.StatusForbidden)
		return

	}
	current := stream.CurrentSize()
	if current == -1 {
		c.Status(http.StatusNotFound)
		return

	}
	c.Header("content-length", fmt.Sprintf("%d", current))
}
