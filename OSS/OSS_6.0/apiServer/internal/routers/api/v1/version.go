package v1

import (
	"github.com/gin-gonic/gin"
)

type Version struct{}

func NewVersion() Version {
	return Version{}
}

func (this Version) Get(c *gin.Context) {}
