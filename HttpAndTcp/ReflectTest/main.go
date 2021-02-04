package main

import (
	jt "ReflectTest/jsonType"
	"errors"
	"fmt"
	"reflect"
)

//使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体tag的值

func main() {
	a := jt.Json{
		Name: "jiangkun",
		Age:  28,
	}
	aTyp := reflect.TypeOf(a)
	if aTyp.Kind() != reflect.Struct {
		err := errors.New("不是结构体")
		panic(err)
	}
	aVal := reflect.ValueOf(&a)
	aVal.Elem().FieldByName("Name").SetString("yixingwei")

	num := aTyp.NumField()
	val := reflect.ValueOf(a)
	for i := 0; i < num; i++ {
		res, _ := aTyp.Field(i).Tag.Lookup("json")
		fmt.Println("res: ", res)
		fmt.Println(val.Field(i))
	}

}
