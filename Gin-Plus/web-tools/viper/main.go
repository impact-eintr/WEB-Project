package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// 设置默认值
	viper.SetDefault("fileDir", "./")

	// 读取配置文件
	viper.SetConfigName("config")         // 配置文件名称(注意没有扩展名)
	viper.SetConfigType("yaml")           // 如果配置文件的名称没有扩展名 需要配置此项
	viper.AddConfigPath("/etc/config/")   // 查找配置文件所在的路径
	viper.AddConfigPath("$HOME/.config/") // 多次调用可以添加多个搜索路径
	viper.AddConfigPath("./config")       //还可以在工作目录中查找配置

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file:%s \n", err))
	}
}
