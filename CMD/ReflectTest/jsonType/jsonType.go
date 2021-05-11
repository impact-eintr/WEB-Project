package jsonType

import (
	"fmt"
)

type Json struct {
	Name string `json:"名字"` //反射机制
	Age  int    `json:"年龄"`
	ID   int    `json:"编号"`
	id   int
}

func (j *Json) Setid(id int) {
	j.id = id
}

func (j *Json) Info() {
	fmt.Println(*j)
}
