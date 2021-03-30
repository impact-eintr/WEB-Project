package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Info struct {
}

type L21 struct {
	R路线编号     string `json:"路线编号"`
	R所在行政区划代码 string `json:"所在行政区划代码"`
	R路线名称     string `json:"路线名称"`
	R起点名称     string `json:"起点名称"`
	R止点名称     string `json:"止点名称"`
	R起点桩号     string `json:"起点桩号"`
	R止点桩号     string `json:"止点桩号"`
	R里程公里     string `json:"里程(公里)"`
	//R技术等级代码      string
	//R技术等级        string
	//R是否一幅高速      string
	R车道数量个 string `json:"车道数量(个)"`
	//R面层类型代码      string
	R面层类型     string `json:"面层类型"`
	R路基宽度米    string `json:"路基宽度(米)"`
	R路面宽度米    string `json:"路面宽度(米)"`
	R面层厚度厘米   string `json:"面层厚度(厘米)"`
	R设计时速公里小时 string `json:"设计时速(公里/小时)"`
	//R修建年度        string
	//R改建年度        string
	//R最近一次修复养护年度  string
	//R断链类型        string
	//R是否城管路段      string
	//R是否断头路段      string
	//R路段收费性质      string
	//R重复路段线路编号    string
	//R重复路段起点桩号    string
	//R重复路段终点桩号    string
	//R养护里程        string
	//R可绿化里程      string
	//R已绿化里程      string
	//R地貌代码        string
	//R地貌汉字        string
	//R涵洞数量个       string
	//R管养单位名称      string
	//R省际出入口       string
	//R国道调整前路线编号   string
	//R是否按干线公路管理接养 string
	//R备注          string
}

func main() {
	db, err := sql.Open("mysql", "root:tyutBigData103@tcp(192.168.23.169:3306)/BigData?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Println("OK!")

	rows, err := db.Query("select `路线编号`,`所在行政区划代码`,`路线名称` ,`起点名称`,`止点名称`,`起点桩号`,`止点桩号`,`里程（公里）`,`车道数量(个)`,`面层类型`,`路基宽度(米)`,`路面宽度(米)`,`面层厚度(厘米)`,`设计时速(公里/小时)` from L21 limit 10")
	if err != nil {
		log.Println(err)
		return
	}

	var roads []string
	for rows.Next() {
		var road L21
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

	for _, val := range roads {
		fmt.Println("查询结果：", val)
	}

}
