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

func RoadQuery(count int) (roads string) {
	db, err := sql.Open("mysql", "root:123456789@tcp(192.168.23.169:3306)/BigData?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Println("成功连接到数据库!")

	rows, err := db.Query("select `路线编号`,`所在行政区划代码`,`路线名称` ,`起点名称`,`止点名称`,`起点桩号`,`止点桩号`,`里程（公里）`,`车道数量(个)`,`面层类型`,`路基宽度(米)`,`路面宽度(米)`,`面层厚度(厘米)`,`设计时速(公里/小时)` from L21 limit ?,1000", count)
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		var road common.L21
		rows.Scan(
			&road.R路线编号,
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
		roads += string(data)
		roads += ","
	}
	roads = roads[:len(roads)-1]
	return
}

func BridgeQuery(count int) (bridges string) {
	db, err := sql.Open("mysql", "root:123456789@tcp(192.168.23.169:3306)/BigData?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Println("成功连接到数据库!")

	rows, err := db.Query("select `桥梁名称`,`桥梁代码`,`桥梁中心桩号` ,`路线编号`,`路线名称`,`技术等级`,`桥梁全长(米)`,`跨径总长（米）`,`单孔最大跨径(米)`,`桥梁组合)孔*米)`,`桥梁全宽(米)`,`桥面净宽(米)`,`按跨径分类代码`,`按跨径分类类型` from L24a limit ?,2", count)
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		var bridge common.L24a
		rows.Scan(&bridge.Q桥梁名称,
			&bridge.Q桥梁代码,
			&bridge.Q桥梁中心桩号,
			&bridge.Q路线编号,
			&bridge.Q路线名称,
			&bridge.Q技术等级,
			&bridge.Q桥梁全长米,
			&bridge.Q跨径总长米,
			&bridge.Q单孔最大跨径米,
			&bridge.Q跨径组合孔米,
			&bridge.Q桥梁全宽米,
			&bridge.Q桥面净宽米,
			&bridge.Q按跨径分类代码,
			&bridge.Q按跨径分类类型)
		data, err := json.Marshal(bridge)
		if err != nil {
			log.Println(err)
		}
		bridges += string(data)
		bridges += ","
	}
	bridges = bridges[:len(bridges)-1]
	return
}

func TunnelQuery(count int) (tunnels string) {
	db, err := sql.Open("mysql", "root:123456789@tcp(192.168.23.169:3306)/BigData?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Println("成功连接到数据库!")

	rows, err := db.Query("select `路线编号`,`所在行政区划代码`,`路线名称` ,`起点名称`,`止点名称`,`起点桩号`,`止点桩号`,`里程（公里）`,`车道数量(个)`,`面层类型`,`路基宽度(米)`,`路面宽度(米)`,`面层厚度(厘米)`,`设计时速(公里/小时)` from L21 limit ?,1000", count)
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		var tunnel common.L25
		rows.Scan(
			&tunnel.S隧道名称,
			&tunnel.S隧道代码,
			&tunnel.S隧道中心桩号,
			&tunnel.S所属路线编号,
			&tunnel.S所属路线名称,
			&tunnel.S所属线路技术等级,
			&tunnel.S隧道长度米,
			&tunnel.S隧道净宽米,
			&tunnel.S隧道净高米,
			&tunnel.S隧道按长度分类代码,
			&tunnel.S隧道按长度分类)
		data, err := json.Marshal(tunnel)
		if err != nil {
			log.Println(err)
		}
		tunnels += string(data)
		tunnels += ","
	}
	tunnels = tunnels[:len(tunnels)-1]
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
	infotype, _ := c.Get("infotype")

	var info string
	switch infotype {
	case "road":
		info = RoadQuery(countnum)
	case "bridge":
		info = BridgeQuery(countnum)
	case "tunnel":
		info = TunnelQuery(countnum)

	}
	c.Set("info", info)
	c.Next()
}

func main() {
	r := gin.Default()

	//typ := flag.String("type", "rocksdb", "cache type")
	typ := flag.String("type", "inmemory", "cache type")
	ttl := flag.Int("ttl", 0, "TTL")
	flag.Parse()
	log.Println("type is", *typ)

	c := cache.New(*typ, *ttl)

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
			info := c.GetString("info")
			key := "/" + c.Param("infotype") + "/" + c.Param("count")

			//var x = []byte{}

			//for i, l := 0, len(info); i < l; i++ {
			//	b := []byte(info[i])
			//	for j := 0; j < len(b); j++ {
			//		x = append(x, b[j])
			//	}
			//}

			c.JSON(http.StatusOK, info)
			c.Request.URL.Path = "/cache/update" + key //将请求的URL修改
			c.Request.Method = http.MethodPut
			c.Request.Body = ioutil.NopCloser(bytes.NewReader([]byte(info)))

			r.HandleContext(c) //继续之后的操作
		})
	}

	r.Run("0.0.0.0:8081")
}
