package router

import (
	"basic/global"
	"basic/internal/dao/webcache/cache"
	"basic/internal/middleware"
	"basic/pkg/tcp"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	v1 "basic/internal/router/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	// 调用gin自带的日志收集 之后可以替换
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 配置缓存服务
	c := cache.New(global.CacheSetting.CacheType, global.CacheSetting.TTL)
	s := v1.NewServer(c)
	// 开启缓存服务
	go tcp.New(c).Listen()

	apiv1 := r.Group("/api/v1")
	{
		// 缓存路由
		{
			cacheGroup := apiv1.Group("/cache")
			cacheGroup.Use(middleware.Cors(), middleware.PathParse)
			cacheGroup.GET("/hit/*key", s.CacheCheck, func(c *gin.Context) {
				miss := c.GetBool("miss") // 检查是否命中缓存
				if miss {
					c.Request.URL.Path = "/info" + c.Param("key") // 将请求的URL修改
					r.HandleContext(c)                            // 继续之后的操作

				}
			})

			cacheGroup.PUT("/update/*key", s.UpdateHandler)
			cacheGroup.DELETE("/delete/*key", s.DeleteHandler)
			cacheGroup.GET("/status/", s.StatusHandler)
		}

		// 数据查询路由
		infoGroup := apiv1.Group("/info")
		{
			infoGroup.Use(middleware.Cors(), middleware.PathParse)
			infoGroup.GET("/:infotype/:count", middleware.QueryRouter, func(c *gin.Context) {
				info := c.GetString("info")

				c.JSON(http.StatusOK, info) // 向浏览器返回数据

				key := "/" + c.Param("infotype") + "/" + c.Param("count")

				// 如果是缓存在磁盘中
				if global.CacheSetting.CacheType == "disk" {
					go func(string, string) {
						klen := strconv.Itoa(len(key))
						vlen := strconv.Itoa(len(info))
						test := "S" + klen + " " + vlen + " " + key + info

						serverAddr := fmt.Sprintf("127.0.0.1:%s", global.CacheSetting.Port)
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

					}(key, info)
				} else {
					c.Request.URL.Path = "/api/v1/cache/update" + key //将请求的URL修改
					c.Request.Method = http.MethodPut
					c.Request.Body = ioutil.NopCloser(bytes.NewReader([]byte(info)))

					r.HandleContext(c) //继续之后的操作
				}
			})
		}

	}
	return r
}
