package main

import (
	"basic/global"
	"basic/internal/dao/webcache/cache"
	"basic/internal/router"
	"basic/pkg/cachehttp"
	"basic/pkg/setting"
	"basic/pkg/tcp"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func init() {
	err := SettingInit()
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {

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

	r.Run("0.0.0.0:8081")
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

	// 实时监控配置文件的变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("修改配置文件", e.Name)

	})
	return nil

}
