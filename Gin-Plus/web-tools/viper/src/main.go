package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var config *viper.Viper

func main() {
	config = initConfigure()
	r := gin.Default()
	r.GET("/getConfig", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"config": config.AllSettings(),
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080

}

func initConfigure() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")    // 设置文件名称（无后缀）
	v.SetConfigType("yaml")      // 设置后缀名 {"1.6以后的版本可以不设置该后缀"}
	v.AddConfigPath("../config") // 设置文件所在路径
	v.Set("verbose", true)       // 设置默认参数

	log.Println(v)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(" Config file not found; ignore error if desired")

		} else {
			panic("Config file was found but another error was produced")

		}

	}
	// 监控配置和重新获取配置
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)

	})
	return v

}
