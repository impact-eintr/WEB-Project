package main

import (
	"OSS/apiServer/global"
	"OSS/apiServer/internal/routers"
	"OSS/apiServer/pkg/setting"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/gin-gonic/gin"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return nil
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func main() {
	fmt.Println(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:         global.ServerSetting.ListenAddr,
		Handler:      router,
		ReadTimeout:  global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
	}

	s.ListenAndServe()
}
