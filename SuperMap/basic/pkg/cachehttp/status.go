package cachehttp

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) StatusHandler(c *gin.Context) {
	log.Println(s.GetStat())
	b, err := json.Marshal(s.GetStat())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, string(b))
}
