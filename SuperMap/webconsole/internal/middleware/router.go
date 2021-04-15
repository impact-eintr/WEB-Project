package middleware

import (
	"strconv"
	"webconsole/internal/dao/list"

	"github.com/gin-gonic/gin"
)

// 获取路径的中间件
func PathParse(c *gin.Context) {
	infotype := c.Param("infotype")
	count := c.Param("count")

	c.Set("infotype", infotype)
	c.Set("count", count)

	c.Next()
}

// 处理路由信息的中间件
func QueryRouter(c *gin.Context) {
	count, _ := c.Get("count")
	countnum, _ := strconv.Atoi(count.(string))
	infotype, _ := c.Get("infotype")

	var info string
	switch infotype {
	case "road":
		info = list.RoadQuery(countnum)
	case "bridge":
		info = list.BridgeQuery(countnum)
	case "tunnel":
		info = list.TunnelQuery(countnum)
	case "service":
		info = list.FQuery(countnum)
	case "portal":
		info = list.MQuery(countnum)
	case "toll":
		info = list.SQuery(countnum)
	}
	c.Set("info", info)
	c.Next()
}
