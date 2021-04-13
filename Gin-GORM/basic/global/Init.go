package global

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func ConfigInit() {
	// 设置默认值
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./confs/")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误
		} else {
			// 配置文件找到后发生了其他错误
		}
	}

	G.Port = viper.GetInt("port")
	G.CacheConfig.CacheType = viper.GetString("cache.cachetype")
	if G.CacheConfig.CacheType != "" {
		G.CacheConfig.Port = viper.GetInt("cache.port")
		G.CacheConfig.TTL = viper.GetInt("cache.ttl")
		G.CacheConfig.CacheDir = viper.GetString("cache.cacheDir")

	} else {
		// 如果不设置缓存，可以直接连接到数据库(待设计)

	}

	G.MysqlConfig.Host = viper.GetString("mysql.host")
	G.MysqlConfig.Port = viper.GetInt("mysql.port")
	G.MysqlConfig.User = viper.GetString("mysql.user")
	G.MysqlConfig.Password = viper.GetString("mysql.password")
	G.MysqlConfig.DBname = viper.GetString("mysql.dbname")

	// 实时监控配置文件的变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("修改配置文件", e.Name)

	})

	log.Println("cache type is", G.CacheConfig.CacheType)

}
