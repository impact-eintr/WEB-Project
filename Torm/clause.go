package Torm

import (
	"fmt"
	"reflect"
	"strings"
)

//Clause 条款、子句
type Clause struct {
	Cselect   string //查询字段
	Cset      string //修改字段
	Tablename string //表名

	Condition string //查询条件
	Limit     int32  //查询条件
	Offset    int32

	Sql    string //完整sql语句
	Params []interface{}

	SqlType    map[Type]string
	ParamsType map[Type][]interface{}
}

func NewClause() *Clause {
	return &Clause{
		Cselect:    "*",
		Limit:      -1,
		Offset:     -1,
		SqlType:    make(map[Type]string),
		ParamsType: make(map[Type][]interface{}),
	}
}

func (this *Clause) SetTableName(tablename string) *Clause {
	this.Tablename = tablename
	return this
}

//根据关键字构建sql语句
func (this *Clause) Set(operation Type, param ...interface{}) {
	sql, vars := generators[operation](param...)
	fmt.Println("vars : ", vars)
	this.SqlType[operation] = sql
	this.ParamsType[operation] = vars
}

//拼接各个sql语句
func (this *Clause) Build(orders ...Type) {
	var sqls []string
	var vars []interface{}
	for _, order := range orders {
		if sql, ok := this.SqlType[order]; ok {
			sqls = append(sqls, sql)
			vars = append(vars, this.ParamsType[order]...)
		}
	}

	this.Sql = strings.Join(sqls, " ")
	this.Params = vars
}

func (this *Clause) InsertStruct(vars interface{}) *Clause {
	typ := reflect.TypeOf(vars)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return this
	}

	//数据映射到
	schema := StructForType(typ)

	//构建SQL语句
	this.Set(Insert, this.Tablename, schema.FieldNames) //["Name","Age"]
	//INSERT INTO test(name,age)
	recordValues := make([]interface{}, 0)
	recordValues = append(recordValues, schema.RecordValues(vars))
	//recordValues := schema.RecordValues(vars)
	this.Set(Value, recordValues...)
	this.Build(Insert, Value)
	return this

}

func (c *Clause) UpdateStruct(vars interface{}) *Clause {
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
	c.Set(Update, c.Tablename, m)
	return c

}

func (c *Clause) AndEqual(field string, value interface{}) *Clause {
	return c.setCondition(Condition, "AND", field, "=", value)

}

func (c *Clause) OrEqual(field string, value interface{}) *Clause {
	return c.setCondition(Condition, "OR", field, "=", value)

}

// 查询字段
func (c *Clause) SelectField(cselect ...string) *Clause {
	c.Cselect = strings.Join(cselect, ",")
	return c

}

// 查询条件组装
func (c *Clause) setCondition(values ...interface{}) *Clause {
	sql, vars := generators[values[0].(Type)](values[2:]...)
	c.Params = append(c.Params, vars...)
	c.addCondition(sql, values[1].(string))
	return c

}

// 条件组成
func (c *Clause) addCondition(sql, opt string) {
	if c.Condition == "" {
		c.Condition = sql

	} else {
		c.Condition = fmt.Sprint("(", c.Condition, ") ", opt, " (", sql, ")")

	}

}
