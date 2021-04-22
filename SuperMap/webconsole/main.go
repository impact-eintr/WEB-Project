package main

import (
	"context"
	"fmt"
	"syscall"
	"time"
	"webconsole/global"

	"go.uber.org/zap"

	"log"
	"net/http"
	"os"
	"os/signal"
	"webconsole/internal/dao/database"
	"webconsole/internal/router"
	"webconsole/pkg/logger"
	"webconsole/pkg/setting"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 加载配置文件
	if err := global.Init(); err != nil {
		fmt.Println("init failed, err: ", err)
		return
	}

	err := global.Conf.ReadSection("server", &global.ServerSetting)
	if err != nil {
		fmt.Println("init failed, err: ", err)
		return
	}

	// 初始化日志
	err = global.Conf.ReadSection("log", &global.LoggerSetting)
	if err != nil {
		fmt.Println("init logger failed, err: ", err)
		return
	}

	if err := logger.Init(); err != nil {
		fmt.Println("init logger failed, err: ", err)
		return
	}

	zap.L().Debug("logger init success...")

	// 初始化缓存
	err = global.Conf.ReadSection("cache", &global.CacheSetting)
	if err != nil {
		fmt.Println("init cache failed, err: ", err)
		return
	}

	if ctyp := global.CacheSetting.CacheType; ctyp != "" {
		zap.L().Debug("cache init success...", zap.String("cachetype", ctyp))
	} else {
		// 如果不设置缓存，可以直接连接到数据库(待设计)
		log.Fatalln("未指定缓存类型")
	}

	zap.L().Debug("cache init success...")

	err = global.Conf.ReadSection("database", &global.DatabaseSetting)
	if err != nil {
		fmt.Println("init database failed, err: ", err)
		return
	}

	// 初始化sql连接
	if err := database.Init(); err != nil {
		fmt.Println("init database failed, err: ", err)
		return
	}

	zap.L().Debug("database init success...")

}

// @title 交通一张图后端系统
// @version 1.0.0
// @description 交通一张图
func main() {

	defer zap.L().Sync()
	r, err := router.NewRouter()
	if err != nil {
		return
	}

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

	err = setting.ReadSection("server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("cache", &global.CacheSetting)
	if err != nil {
		return err
	}

	if ctyp := global.CacheSetting.CacheType; ctyp != "" {
		log.Println("cache type is", ctyp)
	} else {
		// 如果不设置缓存，可以直接连接到数据库(待设计)
		log.Fatalln("未指定缓存类型")
	}

	return nil

}
