package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	sf "bluebell/pkg/snowflake"
	"bluebell/router"
	"bluebell/setting"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	// 1. 加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Printf("init mysql failed, err:#{err}\n")
		return
	}

	// 2. 初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init mysql failed, err:#{err}\n")
		return
	}

	zap.L().Debug("logger init success...")

	defer zap.L().Sync()

	// 3. 初始化MySQL连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err:#{err}\n")
		return
	}
	zap.L().Debug("mysql init success...")
	defer mysql.Close()

	// 4. 初始化缓存
	if err := redis.Init(); err != nil {
		fmt.Printf("init redis failed, err:#{err}\n")
		return
	}
	zap.L().Debug("redis init success...")
	defer redis.Close()

	// 初始化ID生成器
	if err := sf.Init("2021-04-16", setting.Conf.MachineID); err != nil {
		fmt.Printf("init failed,err:%v\n", err)
		return
	}

	if err := controller.InitTrans(setting.Conf.AppConfig.Locale); err != nil {
		fmt.Println(err)
		return
	}

	// 5. 注册路由
	r := router.Setup()

	// 6. 启动服务(优雅关机)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Conf.AppConfig.Port),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			zap.L().Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	// 优雅关机
	quit := make(chan os.Signal, 1) // 创建一个接受信号的信道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞在此处

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal("Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exit")

}
