package v1

import (
	"github.com/gin-gonic/gin"
)

type Temp struct{}

func NewTemp() Temp {
	return Temp{}
}

func (this Temp) Put(c *gin.Context)  {}
func (this Temp) Head(c *gin.Context) {}
