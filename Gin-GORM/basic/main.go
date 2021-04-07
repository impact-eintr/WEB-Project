package main

import (
	"basic/cache/cache"
	cachehttp "basic/cache/http"
	"basic/cache/tcp"
	"basic/common"
	"basic/middleware"
	"bytes"
	"io/ioutil"

	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Query(count int) (roads []string) {
	db, err := sql.Open("mysql", "root:123456789@tcp(192.168.23.169:3306)/BigData?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Println("成功连接到数据库!")

	rows, err := db.Query("select `路线编号`,`所在行政区划代码`,`路线名称` ,`起点名称`,`止点名称`,`起点桩号`,`止点桩号`,`里程（公里）`,`车道数量(个)`,`面层类型`,`路基宽度(米)`,`路面宽度(米)`,`面层厚度(厘米)`,`设计时速(公里/小时)` from L21 limit ?,3", count)
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		var road common.L21
		rows.Scan(&road.R路线编号,
			&road.R所在行政区划代码,
			&road.R路线名称,
			&road.R起点名称,
			&road.R止点名称,
			&road.R起点桩号,
			&road.R止点桩号,
			&road.R里程公里,
			&road.R车道数量个,
			&road.R面层类型,
			&road.R路基宽度米,
			&road.R路面宽度米,
			&road.R面层厚度厘米,
			&road.R设计时速公里小时)
		data, err := json.Marshal(road)
		if err != nil {
			log.Println(err)
		}
		//roads += string(data)
		roads = append(roads, string(data))
	}
	return
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
	roads := Query(countnum)
	c.Set("roads", roads)
	c.Next()
}

func main() {
	r := gin.Default()

	//typ := flag.String("type", "rocksdb", "cache type")
	typ := flag.String("type", "inmemory", "cache type")
	flag.Parse()
	log.Println("type is", *typ)

	c := cache.New(*typ)

	// 开启缓存服务
	go tcp.New(c).Listen()

	cacheGroup := r.Group("/cache")
	{
		cacheGroup.Use(middleware.Cors(), m1)

		cacheGroup.Any("/hit/*key", cachehttp.New(c).CacheCheck, func(c *gin.Context) {
			test, _ := c.Get("test")
			if test.(bool) {
				c.Request.URL.Path = "/info" + c.Param("key") //将请求的URL修改
				r.HandleContext(c)                            //继续之后的操作
			}
		})

		cacheGroup.PUT("/update/*key", cachehttp.New(c).UpdateHandler)

		cacheGroup.GET("/status/", cachehttp.New(c).StatusHandler)
	}

	infoGroup := r.Group("/info")
	{
		infoGroup.Use(middleware.Cors(), m1)

		infoGroup.GET("/:infotype/:count", m3, func(c *gin.Context) {
			roads := c.GetStringSlice("roads")
			//roads := c.GetString("roads")
			key := "/" + c.Param("infotype") + "/" + c.Param("count")

			var x = []byte{}

			for i, l := 0, len(roads); i < l; i++ {
				b := []byte(roads[i])
				for j := 0; j < len(b); j++ {
					x = append(x, b[j])
				}
			}

			c.JSON(http.StatusOK, roads)
			c.Request.URL.Path = "/cache/update" + key //将请求的URL修改
			c.Request.Method = http.MethodPut
			c.Request.Body = ioutil.NopCloser(bytes.NewReader(x))

			r.HandleContext(c) //继续之后的操作

		})
	}

	r.Run(":8081")
}
