package Torm_test

import (
	"Torm"
	log "github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

type User struct {
	Name string `torm:"user_name,varchar"`
	Age  int    `torm:"age,int"`
}

func TestStructForType(t *testing.T) {
	user := &User{}
	utypes := reflect.TypeOf(user)
	schema := Torm.StructForType(utypes.Elem()) //Elem取出指针指向的值
	log.Info(schema.FieldNames)
	for i := 0; i < len(schema.Fields); i++ {
		log.Info("字段名称：", schema.Fields[i].Name, ";字段类型:", schema.Fields[i].Type,
			";对应数据库列名:", schema.Fields[i].TableColumn)

	}
	if len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")

	}
	if schema.FieldMap["Name"].Name != "Name" {
		t.Fatal("failed to parse primary key")

	}

}
