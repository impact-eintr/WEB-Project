package database

import (
	"database/sql"
	"fmt"
	"webconsole/global"

	_ "github.com/go-sql-driver/mysql"
)

func Init() error {
	var err error
	dbinfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.DatabaseSetting.User,
		global.DatabaseSetting.Password,
		global.DatabaseSetting.Host,
		global.DatabaseSetting.Port,
		global.DatabaseSetting.DBname,
	)

	global.DB, err = sql.Open("mysql", dbinfo)
	if err != nil {
		return err
	}

	err = global.DB.Ping()
	if err != nil {
		return err
	}

	// 根据具体需求设置
	//global.DB.SetConnMaxIdleTime(time.Second * 10)
	//global.DB.SetMaxOpenConns(200)
	//global.DB.SetMaxIdleConns(10)

	return nil
}
