package middleware

import (
	"strconv"
	"webconsole/internal/dao/database"

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
		info = database.RoadQuery(countnum)
	case "bridge":
		info = database.BridgeQuery(countnum)
	case "tunnel":
		info = database.TunnelQuery(countnum)
	case "service":
		info = database.FQuery(countnum)
	case "portal":
		info = database.MQuery(countnum)
	case "toll":
		info = database.SQuery(countnum)
	}
	c.Set("info", info)
	c.Next()
}
