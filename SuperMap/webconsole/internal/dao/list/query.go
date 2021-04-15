package list

import (
	"encoding/json"
	"log"
	"webconsole/global"
	"webconsole/internal/model"

	_ "github.com/go-sql-driver/mysql"
)

func Level(level int) string {
	switch level {
	case 0:
		return "高速"
	case 1:
		return "一级"
	case 2:
		return "二级"
	case 3:
		return "三级"
	case 4:
		return "四级"
	case 5:
		return "等外"
	}
	return ""
}

func RoadQuery(count int) string {
	level := Level(count)

	rows, err := global.DB.Query("select `路线编号`,`所在行政区划代码`,`路线名称` ,`起点名称`,`止点名称`,`起点桩号`,`止点桩号`,`里程（公里）`,`车道数量(个)`,`面层类型`,`路基宽度(米)`,`路面宽度(米)`,`面层厚度(厘米)`,`设计时速(公里/小时)` from L21 where `技术等级`=? AND `ID`>2", level)
	if err != nil {
		log.Println(err)
		return ""
	}

	roads := []model.L21{}

	for rows.Next() {
		var road model.L21
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
		roads = append(roads, road)
	}

	data, err := json.Marshal(roads)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func BridgeQuery(count int) string {
	level := Level(count)

	rows, err := global.DB.Query("select `桥梁名称`, `桥梁代码`, `桥梁中心桩号`, `路线编号`, `路线名称`, `技术等级`, `桥梁全长(米)`, `跨径总长(米)`, `单孔最大跨径(米)`, `跨径组合(孔*米)`, `桥梁全宽(米)`, `桥面净宽(米)`, `按跨径分类代码`, `按跨径分类类型` from L24a where `技术等级`=? AND `ID`>2", level)
	if err != nil {
		log.Println(err)
		return ""
	}

	bridges := []model.L24a{}

	for rows.Next() {
		var bridge model.L24a
		rows.Scan(
			&bridge.Q桥梁名称,
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
		bridges = append(bridges, bridge)
	}

	data, err := json.Marshal(bridges)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func TunnelQuery(count int) string {
	level := Level(count)

	rows, err := global.DB.Query("select `隧道名称`,`隧道代码`,`隧道中心桩号` ,`所属路线编号`,`所属路线名称`,`所属线路技术等级`,`隧道长度(米)`,`隧道净宽(米)`,`隧道净高(米)`,`隧道按长度分类代码`,`隧道按长度分类` from L25 where `所属线路技术等级`=? AND `ID`>2", level)
	if err != nil {
		log.Println(err)
		return ""
	}

	tunnels := []model.L25{}

	for rows.Next() {
		var tunnel model.L25
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
		tunnels = append(tunnels, tunnel)
	}
	data, err := json.Marshal(tunnels)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func FQuery(count int) string {
	rows, err := global.DB.Query("select `路线编号`,`路线名称`,`桩号` ,`服务设施类型`,`服务设施名称`,`初始运营时间`,`布局形式`,`经营模式`,`占地面积(平方米)`,`停车场面积(平方米)`,`停车位数量(个)` from F where `ID`>2")
	if err != nil {
		log.Println(err)
		return ""
	}

	services := []model.F{}

	for rows.Next() {
		var f model.F
		rows.Scan(
			&f.F路线编号,
			&f.F路线名称,
			&f.F桩号,
			&f.F服务设施类型,
			&f.F服务设施名称,
			&f.F初始运营时间,
			&f.F布局形式,
			&f.F经营模式,
			&f.F占地面积,
			&f.F停车场面积,
			&f.F停车位数量)
		services = append(services, f)
	}
	data, err := json.Marshal(services)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func MQuery(count int) string {
	rows, err := global.DB.Query("select `序号`, `收费门架编号`, `门架名称`, `门架类型`, `门架种类`, `门架标志`, `省界入出口标识`, `收费单元编码组合`, `车道数`, `纬度`, `经度`, `桩号`, `使用状态` from SM ")
	if err != nil {
		log.Println(err)
		return ""
	}

	portals := []model.SM{}

	for rows.Next() {
		var portal model.SM
		rows.Scan(
			&portal.M序号,
			&portal.M收费门架编号,
			&portal.M门架名称,
			&portal.M门架类型,
			&portal.M门架种类,
			&portal.M门架标志,
			&portal.M省界入出口标识,
			&portal.M收费单元编码组合,
			&portal.M车道数,
			&portal.M纬度,
			&portal.M经度,
			&portal.M桩号,
			&portal.M使用状态)
		portals = append(portals, portal)
	}
	data, err := json.Marshal(portals)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func SQuery(count int) string {
	rows, err := global.DB.Query("select `序号`,`收费站编号`,`收费站名称` ,`收费广场数量`,`收费站HEX`,`线路类型`,`网络所属运营商`,`数据汇聚点` from SZ")
	if err != nil {
		log.Println(err)
		return ""
	}

	tolls := []model.SZ{}

	for rows.Next() {
		var toll model.SZ
		rows.Scan(
			&toll.S序号,
			&toll.S收费站编号,
			&toll.S收费站名称,
			&toll.S收费广场数量,
			&toll.S收费站HEX,
			&toll.S线路类型,
			&toll.S网络所属运营商,
			&toll.S数据汇聚点)
		tolls = append(tolls, toll)
	}
	data, err := json.Marshal(tolls)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}
