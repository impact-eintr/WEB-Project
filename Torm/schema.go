package Torm

import (
	"go/ast"
	"reflect"
	"strings"
	"sync"
)

type Field struct {
	Name        string
	i           int
	Type        string
	TableColumn string
	Tag         string
}

type Schema struct {
	Fields     []*Field          //字段属性组合
	FieldNames []string          //字段名称
	FieldMap   map[string]*Field // key:value
}

var (
	structMutex sync.RWMutex
	structCache = make(map[reflect.Type]*Schema)
)

func StructForType(t reflect.Type) *Schema {
	//从缓存中获取 先看可读数据
	structMutex.RLock()
	st, found := structCache[t]
	structMutex.RUnlock()

	if found {
		return st
	}
	//可读数据中没有，尝试抢占锁，再次尝试获取
	structMutex.Lock()
	defer structMutex.Unlock()
	st, found = structCache[t]
	if found {
		return st
	}
	//缓存中没有 创建数据并缓存
	st = &Schema{
		FieldMap: make(map[string]*Field),
	}
	dataTypeOf(t, st)

	structCache[t] = st
	return st
}

//对象与结构的转换
func dataTypeOf(types reflect.Type, schema *Schema) {
	for i := 0; i < types.NumField(); i++ {
		p := types.Field(i)

		if p.Anonymous || !ast.IsExported(p.Name) {
			continue
		}
		field := &Field{
			Name:        p.Name, //默认是结构体原字段
			i:           i,
			TableColumn: p.Name, //默认是结构体原字段
			Tag:         p.Name, //默认是结构体原字段
		}
		//追加标签
		if tg, ok := p.Tag.Lookup("torm"); ok {
			field.Tag = tg
		}

		//切分标签
		tagArr := strings.Split(field.Tag, ",")
		// 数据库中对应列表名称
		if len(tagArr[0]) > 0 {
			field.TableColumn = tagArr[0]
		}

		// 数据库中对应列表类型
		if len(tagArr) > 1 && len(tagArr[1]) > 0 {
			field.Type = tagArr[1]

		}

		schema.Fields = append(schema.Fields, field)
		schema.FieldMap[p.Name] = field
		schema.FieldNames = append(schema.FieldNames, p.Name)
	}
}

func (s *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range s.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())

	}
	return fieldValues

}

func (s *Schema) UpdateParam(dest interface{}) map[string]interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	m := make(map[string]interface{})
	for _, field := range s.Fields {
		m[field.TableColumn] = destValue.FieldByName(field.Name).Interface()

	}
	return m

}
