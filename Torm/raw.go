package Torm

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

// 定义统一接口 便于后续支持多个数据引擎
type CommonDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)

//CUDR
func (this *Session) Exec() (result sql.Result, err error) {
	defer this.Clear()
	log.Println(this.sql.String(), this.sqlValues)
	if result, err = this.db.Exec(this.sql.String(), this.sqlValues...); err != nil {
		log.Println(err)
	}
	return
}

// 单条查询操作
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Println(s.sql.String(), s.sqlValues)
	return s.DB().QueryRow(s.sql.String(), s.sqlValues...)

}

// 多条查询操作
func (s *Session) Query() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Println(s.sql.String(), s.sqlValues)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlValues...); err != nil {
		log.Println(err)
	}
	return

}

type Session struct {
	db        *sql.DB         //数据库引擎
	tx        *sql.Tx         //数据库事务
	sqlValues []interface{}   //SQL 动态参数
	sql       strings.Builder //SQL 语句
}

func (this *Session) DB() CommonDB {
	if this.tx != nil {
		return this.tx
	}
	return this.db
}

//实例化Session
func NewSession(db *sql.DB) *Session {
	return &Session{db: db}
}

func (this *Session) Clear() {
	this.sql.Reset()
	this.sqlValues = nil
}

func (this *Session) Raw(sql string, values ...interface{}) *Session {
	this.sql.WriteString(sql)
	this.sql.WriteString(" ")
	this.sqlValues = append(this.sqlValues, values...)
	return this
}
