package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/setting"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

func Setup() *gin.Engine {
	if setting.Conf.AppConfig.Mode == "release" {
		gin.SetMode(gin.ReleaseMode) // 设置为发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/signup", controller.SignUpHandler)

	r.POST("/login", controller.LoginHandler)

	return r
}

func translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New())
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		// 修改gin框架中的Validator引擎属性，实现自定制
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "en":
				_ = enTranslations.RegisterDefaultTranslations(v, trans)
			case "zh":
				_ = zhTranslations.RegisterDefaultTranslations(v, trans)
			default:
				_ = enTranslations.RegisterDefaultTranslations(v, trans)
			}

			// 注册翻译器
			c.Set("trans", trans)
		}
		c.Next()
	}
}
