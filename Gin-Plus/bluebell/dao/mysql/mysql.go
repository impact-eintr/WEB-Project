package mysql

import (
	"bluebell/setting"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		setting.Conf.MySQLConfig.User,
		setting.Conf.MySQLConfig.Password,
		setting.Conf.MySQLConfig.Host,
		setting.Conf.MySQLConfig.Port,
		setting.Conf.MySQLConfig.DB,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SetMaxOpenConns(setting.Conf.MySQLConfig.MaxOpenConns)
	db.SetMaxIdleConns(setting.Conf.MySQLConfig.MaxIdleConns)
	return

}

func Close() {
	_ = db.Close()
}
