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

## web框架模式
#### MVC模式

- Model(模型)  模型代表一个存储数据的对象或者JAVA POJO。 它也可以带有逻辑，在数据变化时更新控制器。
- View(视图)视图代表模型包含的数据的可视化
- Controller(控制器) 控制器作用于模型和视图上。它控制数据流向模型对象，并在数据变化时更新视图。它使试图与模型分离开

### CLD模型
- Controller 服务的入口。负责处理路由、参数校验、请求转发
- Logic/Service 逻辑服务层 负责处理业务逻辑
- DAO/Responsitory 负责数据与存储相关功能

### 
