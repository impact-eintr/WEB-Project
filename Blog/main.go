package main

import (
	"Blog/global"
	"Blog/internal/model"
	"Blog/internal/routers"
	"Blog/pkg/logger"
	"Blog/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//将配置文件内容映射到应用配置结构体中
func setupSeting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return nil
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return nil
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil

}

//初始化日志管理
// 不是永久返回 nil ??
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" +
		global.AppSetting.LogFileName +
		global.AppSetting.LogFileExt

	global.Logger = logger.NewLogger(
		&lumberjack.Logger{
			Filename:  fileName,
			MaxSize:   600,
			MaxAge:    10,
			LocalTime: true,
		}, "", log.LstdFlags,
	).WithCaller(2)

	return nil
}

//初始化数据库
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting) //注意这里不可以使用 `:=` 赋值
	if err != nil {
		return err
	}
	return nil
}

func init() {
	err := setupSeting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %s", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %s", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
