# Go Web开发常用组件

## zap日志库
~~~ go
package main

import (
	"net/http"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func main() {
	defer logger.Sync() // 将缓存中的内容刷入磁盘
	HttpGet("http://www.google.com")
	HttpGet("www.google.com")

}

func HttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info(
			"Success...",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
~~~

> 将日志文件写入文件而不是终端


~~~ go
func New(core zapcore.Core, option ...Option) *Logger
~~~
`zapcore.Core` 需要三个配置
1. Encoder(怎么写)

~~~ go
zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
~~~

2. WriteSyncer(写到哪里去)

~~~ go
file, err := os.Create("./test.log")
if err != nil {
	logger.Error("create file fault...", zap.Error(err))
}
~~~

3. LogLevel

~~~ go
package main

import (
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger = zap.New(core)
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, err := os.Create("./test.log")
	if err != nil {
		logger.Error("create file fault...", zap.Error(err))
	}
	return zapcore.AddSync(file)
}

func main() {
	InitLogger()
	defer logger.Sync() // 将缓存中的内容刷入磁盘
	HttpGet("http://www.google.com")
	HttpGet("www.google.com")

}

func HttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info(
			"Success...",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
~~~

> zap本身不提供日志切割的功能
> get get -u github.com/natefinch/lumberjack


~~~ go
func getLogWriter() zapcore.WriteSyncer {
	//file, err := os.OpenFile("./test.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0744)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		Maxsize:    10,
		MaxBackups: 5,
		MaxAge:     2,
		Compress:   false, //是否压缩
	}

	return zapcore.AddSync(lumberJackLogger)
}
~~~

## viper处理配置文件

~~~ yaml
port: 8081

mysql:
  host: "127.0.0.1"
  port: 3306
  user: "root"
  password: "1234567"
  dbname: "test"

memcache:
  port: 9425
  ttl: 0

diskcache:
  port: 9425
  ttl: 0

~~~

~~~ go
func main() {
	// 设置默认值
	viper.SetDefault("ttl", 0)

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

	// 实时监控配置文件的变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("修改配置文件", e.Name)
	})

	r := gin.Default()
    // 相关逻辑
	r.Run("0.0.0.0:8081")
}
~~~
