package v1

import (
	"github.com/gin-gonic/gin"
)

type Objects struct{}

func NewObjects() Objects {
	return Objects{}
}

func (this Objects) Put(c *gin.Context)    {}
func (this Objects) Post(c *gin.Context)   {}
func (this Objects) Get(c *gin.Context)    {}
func (this Objects) Delete(c *gin.Context) {}
