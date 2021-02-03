package v1

import (
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (this *Article) Get(c *gin.Context)    {}
func (this *Article) List(c *gin.Context)   {}
func (this *Article) Create(c *gin.Context) {}
func (this *Article) Update(c *gin.Context) {}
func (this *Article) Delete(c *gin.Context) {}
