package tmpl

import (
	"OSS/apiServer/es"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//遇事不决写注释
func Get(c *gin.Context) {
	//渲染模板
	info := make(map[string]interface{}, 1)
	name := c.Param("uid")
	files := es.Test()

	info["name"] = name
	info["files"] = files
	log.Println(info)

	c.HTML(http.StatusOK, "index.html", info)
}
