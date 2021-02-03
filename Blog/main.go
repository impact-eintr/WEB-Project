package main

import (
	"Blog/global"
	"Blog/internal/routers"
	"Blog/pkg/setting"
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

func init() {
	err := setupSeting()
	if err != nil {
		log.Println(err)
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
