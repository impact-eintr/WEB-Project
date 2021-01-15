package main

import "encoding/json"
import "fmt"
import "log"
import "os"

//序列化
//对于结构体的序列化如果我们希望序列化后的key名称，可以自己定制可以像下面这样给属性加一个tag
type Json struct {
	Name string `json:"名字"` //反射机制
	Age  int    `json:"年龄"`
}

func testStuct() {
	a := Json{
		Name: "jiangkun",
		Age:  28,
	}
	res, err := json.Marshal(a)
	if err != nil {
		log.Println(err)
	}
	os.Stdout.Write(res)
}

func testMap() []byte {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "姜昆"
	a["age"] = 28
	data, err := json.Marshal(a)
	if err != nil {
		log.Fatalln(err)
	}
	os.Stdout.Write(data)
	return data
}

func testSlince() []byte {
	var Slince []map[string]Json
	bigdata := make(map[string]Json)

	bigdata["导员"] = Json{
		Name: "姜昆",
		Age:  28,
	}
	bigdata["1840"] = Json{
		Name: "宋NB",
		Age:  22,
	}

	明泽苑 := make(map[string]Json)
	明泽苑["2-103"] = Json{
		Name: "宋NB",
		Age:  22,
	}

	Slince = append(Slince, bigdata, 明泽苑)

	data, err := json.Marshal(Slince)
	if err != nil {
		log.Println(err)
	}
	//os.Stdout.Write(data)
	return data

}

//反序列化
func Slincetest(src []byte) {
	var Slince []map[string]Json

	err := json.Unmarshal(src, &Slince)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(Slince)
}

func main() {
	Slincetest(testSlince())
}
