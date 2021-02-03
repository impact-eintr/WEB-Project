package v1

import (
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (this *Tag) Get(c *gin.Context)    {}
func (this *Tag) List(c *gin.Context)   {}
func (this *Tag) Create(c *gin.Context) {}
func (this *Tag) Update(c *gin.Context) {}
func (this *Tag) Delete(c *gin.Context) {}
