package main

import (
	"basic/cache/cache"
	"basic/cache/tcp"
	"basic/common"
	"basic/middleware"
	"flag"

	"database/sql"
	"encoding/json"
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

	log.Println("OK!")

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
		roads = append(roads, string(data))
	}
	return
}

func m1(c *gin.Context) {
	num, _ := strconv.Atoi(c.Param("id"))
	log.Println(num)
	roads := Query(num)
	c.Set("roads", roads)
	c.Next()

}

func m2(c *gin.Context) {
	num, _ := strconv.Atoi(c.Param("id"))
	log.Println(num)
	roads := Query(num)
	c.Set("roads", roads)
	c.Next()

}

func main() {

	typ := flag.String("type", "inmemory", "cache type")
	flag.Parse()
	log.Println("type is ", *typ)
	c := cache.New(*typ)
	go tcp.New(c).Listen()
	//cachehttp.New(c).Listen()

	r := gin.Default()

	r.Use(middleware.Cors())

	r.GET("/road-info/:id", m1, func(c *gin.Context) {
		roads, err := c.Get("roads")
		if !err {
			log.Fatalln(err)
		}
		data := roads
		c.JSON(http.StatusOK, data)
	})

	key := strings.Split(r.URL.EscapedPath(), "/")[2]
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	m := r.Method
	if m == http.MethodPut {
		b, _ := io.ReadAll(r.Body)
		if len(b) != 0 {
			e := h.Set(key, b)
			if e != nil {
				log.Println(e)
				w.WriteHeader(http.StatusInternalServerError)

			}

		}
		return

	}
	if m == http.MethodGet {
		b, e := h.Get(key)
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		if len(b) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return

		}
		w.Write(b)
		return

	}
	if m == http.MethodDelete {
		e := h.Del(key)
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)

		}
		return

	}

	r.Run(":8081")
}
