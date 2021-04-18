package main

import (
	"context"
	"database/sql"
	"fmt"
	"syscall"
	"time"
	"webconsole/global"

	"log"
	"net/http"
	"os"
	"os/signal"
	"webconsole/internal/router"
	"webconsole/pkg/setting"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 初始化各种配置
	err := SettingInit()
	if err != nil {
		log.Fatalln(err)
	}

	// 初始化sql连接
	err = DBInit()
	if err != nil {
		log.Fatalln(err)
	}

}

// @title 交通一张图后端系统
// @version 1.0.0
// @description 交通一张图
func main() {
	r := router.NewRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.ServerSetting.Port),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 优雅关机
	quit := make(chan os.Signal, 1) // 创建一个接受信号的信道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞在此处

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 延时关闭数据库连接(可能有坑)
	defer global.DB.Close()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("Shutdown", err)
	}

	log.Println("Server exit")

}

func SettingInit() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Cache", &global.CacheSetting)
	if err != nil {
		return err
	}

	if ctyp := global.CacheSetting.CacheType; ctyp != "" {
		log.Println("cache type is", ctyp)
	} else {
		// 如果不设置缓存，可以直接连接到数据库(待设计)
		log.Fatalln("未指定缓存类型")
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		log.Fatalln(err)
	}

	return nil

}

func DBInit() error {
	dbinfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.DatabaseSetting.User,
		global.DatabaseSetting.Password,
		global.DatabaseSetting.Host,
		global.DatabaseSetting.Port,
		global.DatabaseSetting.DBname,
	)

	var err error
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

	log.Println("成功连接到数据库!")
	return nil
}
