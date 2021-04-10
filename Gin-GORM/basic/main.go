package main

import (
	"basic/internal/dao/cache/cache"
	cachehttp "basic/internal/dao/cache/http"
	"basic/internal/dao/cache/tcp"
	"basic/internal/dao/list"
	"basic/internal/middleware"
	"fmt"
	"net"
	"os"

	"flag"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {

	// 设置默认值
	viper.SetDefault("ttl", 0)

	r := gin.Default()

	typ := flag.String("type", "rocksdb", "cache type")
	//typ := flag.String("type", "inmemory", "cache type")
	ttl := flag.Int("ttl", 0, "TTL")
	flag.Parse()
	log.Println("type is", *typ)

	//c := cache.New(*typ, *ttl)
	c := cache.New(*typ, *ttl)

	// 开启缓存服务
	go tcp.New(c).Listen()

	cacheGroup := r.Group("/cache")
	{
		cacheGroup.Use(middleware.Cors(), m1)
		cacheGroup.Any("/hit/*key", cachehttp.New(c).CacheCheck, func(c *gin.Context) {
			miss, _ := c.Get("miss") // 检查是否命中缓存
			if miss.(bool) {
				c.Request.URL.Path = "/info" + c.Param("key") // 将请求的URL修改
				r.HandleContext(c)                            // 继续之后的操作
			}
		})

		cacheGroup.PUT("/update/*key", cachehttp.New(c).UpdateHandler)
		cacheGroup.GET("/status/", cachehttp.New(c).StatusHandler)
	}

	infoGroup := r.Group("/info")
	{
		infoGroup.Use(middleware.Cors(), m1)
		infoGroup.GET("/:infotype/:count", m3, func(c *gin.Context) {
			info := c.GetString("info")
			key := "/" + c.Param("infotype") + "/" + c.Param("count")

			klen := strconv.Itoa(len(key))
			vlen := strconv.Itoa(len(info))
			test := "S" + klen + " " + vlen + " " + key + info

			serverAddr := "127.0.0.1:9425"
			tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
			if err != nil {
				log.Println(err.Error())
				os.Exit(1)
			}

			conn, err := net.DialTCP("tcp", nil, tcpAddr)
			defer conn.Close()

			if err != nil {
				log.Println(err.Error())
				os.Exit(1)

			}
			_, err = conn.Write([]byte(test))
			if err != nil {
				log.Println("Write to server failed:", err.Error())
				os.Exit(1)

			}

			reply := make([]byte, 1024)
			_, err = conn.Read(reply)
			if err != nil {
				log.Println("Write to server failed:", err.Error())
				os.Exit(1)
			}
			fmt.Printf("reply from server:\n%v\n", string(reply))

			//c.JSON(http.StatusOK, info)
			//c.Request.URL.Path = "/cache/update" + key //将请求的URL修改
			//c.Request.Method = http.MethodPut
			//c.Request.Body = ioutil.NopCloser(bytes.NewReader([]byte(info)))

			//r.HandleContext(c) //继续之后的操作
		})
	}

	r.Run("0.0.0.0:8081")
}

// 获取路径的中间件
func m1(c *gin.Context) {
	infotype := c.Param("infotype")
	count := c.Param("count")

	c.Set("infotype", infotype)
	c.Set("count", count)

	c.Next()
}

func m3(c *gin.Context) {
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
