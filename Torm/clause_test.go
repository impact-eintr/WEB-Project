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

//func TestClause_Condition(t *testing.T) {
//	clause := newClause()
//	clause = clause.SetTableName("memo").
//		andEqual("name", "迈莫coding").
//		orEqual("age", 5).
//		selectField("name,age")
//	log.Info(clause.condition)
//	log.Info(clause.params)
//	log.Info(clause.cselect)
//
//}
//func TestClause_UpdateStruct(t *testing.T) {
//	user := &Users{
//		Name: "迈莫coding",
//	}
//	clause := newClause()
//	clause = clause.SetTableName("memo").
//		updateStruct(user)
//	log.Info(clause.sqlType[Update])
//	log.Info(clause.paramsType[Update])
//
//}
