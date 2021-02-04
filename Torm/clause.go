package Torm

import (
	"fmt"
	"reflect"
	"strings"
)

//Clause 条款、子句
type Clause struct {
	cselect   string //查询字段
	cset      string //修改字段
	tablename string //表名

	condition string //查询条件
	limit     int32  //查询条件
	offset    int32

	sql    string //完整sql语句
	params []interface{}

	sqlType    map[Type]string
	paramsType map[Type][]interface{}
}

func NewClause() *Clause {
	return &Clause{
		cselect:    "*",
		limit:      -1,
		offset:     -1,
		sqlType:    make(map[Type]string),
		paramsType: make(map[Type][]interface{}),
	}
}

func (this *Clause) SetTableName(tablename string) *Clause {
	this.tablename = tablename
	return this
}

//根据关键字构建sql语句
func (this *Clause) Set(operation Type, param ...interface{}) {
	sql, vars := generators[operation](param...)
	this.sqlType[operation] = sql
	this.paramsType[operation] = vars
}

//拼接各个sql语句
func (this *Clause) Build(orders ...Type) {
	var sqls []string
	var vars []interface{}
	for _, order := range orders {
		if sql, ok := this.sqlType[order]; ok {
			sqls = append(sqls, sql)
			vars = append(vars, this.paramsType[order]...)
		}
	}

	fmt.Println(sqls, "\n", vars)
	this.sql = strings.Join(sqls, " ")
	this.params = vars
}

func (this *Clause) insertStruct(vars interface{}) *Clause {
	typ := reflect.TypeOf(vars)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return this

	}

	//数据映射
	schema := StructForType(typ)

	//构建SQL语句
	this.Set(Insert, this.tablename, schema.FieldNames)
	recordValues := make([]interface{}, 0)
	recordValues = append(recordValues, schema.RecordValues(vars))
	this.Set(Value, recordValues...)
	this.Build(Insert, Value)
	return this

}

func (c *Clause) updateStruct(vars interface{}) *Clause {
	types := reflect.TypeOf(vars)
	if types.Kind() == reflect.Ptr {
		types = types.Elem()

	}
	if types.Kind() != reflect.Struct {
		return c

	}
	// 数据映射
	schema := StructForType(types)
	m := make(map[string]interface{})
	m = schema.UpdateParam(vars)
	// 构建SQL语句
	c.Set(Update, c.tablename, m)
	return c

}

func (c *Clause) andEqual(field string, value interface{}) *Clause {
	return c.setCondition(Condition, "AND", field, "=", value)

}

func (c *Clause) orEqual(field string, value interface{}) *Clause {
	return c.setCondition(Condition, "OR", field, "=", value)

}

// 查询字段
func (c *Clause) selectField(cselect ...string) *Clause {
	c.cselect = strings.Join(cselect, ",")
	return c

}

// 查询条件组装
func (c *Clause) setCondition(values ...interface{}) *Clause {
	sql, vars := generators[values[0].(Type)](values[2:]...)
	c.params = append(c.params, vars...)
	c.addCondition(sql, values[1].(string))
	return c

}

// 条件组成
func (c *Clause) addCondition(sql, opt string) {
	if c.condition == "" {
		c.condition = sql

	} else {
		c.condition = fmt.Sprint("(", c.condition, ") ", opt, " (", sql, ")")

	}

}
