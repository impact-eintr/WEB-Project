package Torm

import (
	"go/ast"
	"reflect"
	"strings"
	"sync"
)

type Filed struct {
	Name       string `json:"name"`
	i          string
	Type       string `json:"type"`
	TableColum string `json:"tablecolum"`
	Tag        string `json:"tag"`
}

type Schema struct {
	Filed      []*Filed          //字段属性组合
	FiledNames []string          //字段名称
	FiledMap   map[string]*Filed // key:value
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
		FiledMap: make(map[string]*Filed),
	}
	dataTypeof(t, st)

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
			Name: p.Name,
			i:    i,
		}

		var tag = field.Name
		field.TableColum = field.Name
		if tg, ok := p.Tag.Lookup("Torm"); ok {
			tag = tg
		}

		tagArr := strings.Split(tag, ",")
		if len(tagArr) > 0 {
			if tagArr[0] == "-" {
				continue
			}
		}
	}
}
