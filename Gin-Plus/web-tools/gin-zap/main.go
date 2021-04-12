package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func main() {
	r := gin.Default()
	r.GET("hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "gin",
		})
	})
	r.Run()
}

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
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     2,
		Compress:   false,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)

		logger.Info(path,
		zap.Int("status", c.Writer.Status()),
		zap.String("method", c.Request.Method)),
		zap.String("path", path)),
		zap.String("query", query)),
		zap.String("ip", c.ClientIP())),
		zap.String("user-agent", c.Request.UserAgent())),
		)
	}
}
