package Torm_test

import (
	"Torm"
	"log"
	"testing"
)

type Users struct {
	Name string `torm:"user_name,varchar"`
	Age  int    `torm:"age,int"`
}

func TestClause_InsertStruct(t *testing.T) {
	user := &Users{
		Name: "yixingwei",
		Age:  23,
	}
	clause := Torm.NewClause()
	clause = clause.SetTableName("test").
		InsertStruct(user)
	log.Println(clause.Sql)
	log.Println(clause.Params)
	// sql := "INSERT INTO memo (Name,Age) VALUES (?,?)"

}

func TestClause_Condition(t *testing.T) {
	clause := Torm.NewClause()
	clause = clause.SetTableName("test").
		AndEqual("user_name", "迈莫coding").
		OrEqual("age", 5).
		SelectField("user_name,age")
	log.Println(clause.Condition)
	log.Println(clause.Params)
	log.Println(clause.Cselect)

}

func TestClause_UpdateStruct(t *testing.T) {
	user := &Users{
		Name: "迈莫coding",
	}
	clause := Torm.NewClause()
	clause = clause.SetTableName("memo").
		UpdateStruct(user)
	log.Println(clause.SqlType[Torm.Update])
	log.Println(clause.ParamsType[Torm.Update])

}
