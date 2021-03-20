package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&UserInfo{})

	u2 := UserInfo{
		ID:     1,
		Name:   "eintr",
		Gender: "male",
		Hobby:  "code",
	}

	db.Create(&u2)
	log.Println("OK!")

}
