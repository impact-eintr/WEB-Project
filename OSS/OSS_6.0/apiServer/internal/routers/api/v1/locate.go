package v1

import (
	"github.com/gin-gonic/gin"
)

type Locate struct{}

func NewLocate() Locate {
	return Locate{}
}

func (this Locate) Get(c *gin.Context) {}
