package main

import (
	"basic/global"
	"basic/internal/dao/webcache/cache"
	"context"
	"syscall"
	"time"

	"basic/internal/router"
	"basic/pkg/cachehttp"
	"basic/pkg/setting"
	"basic/pkg/tcp"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	err := SettingInit()
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {

	log.Println("main() 开始")
	// 配置缓存服务
	c := cache.New(global.CacheSetting.CacheType, global.CacheSetting.TTL)
	s := cachehttp.New(c)

	// 开启缓存服务
	go tcp.New(c).Listen()

	r := gin.Default()

	// 缓存路由组
	router.CacheRoute(r, s)
	// 数据查询路由组
	router.InfoRoute(r)

	server := &http.Server{
		Addr:    ":8081",
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
