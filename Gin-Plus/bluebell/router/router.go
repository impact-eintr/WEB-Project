package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/setting"
	"net/http"

	_ "github.com/gin-contrib/cors"
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

	r.Use(logger.GinLogger(), logger.GinRecovery(true), Cors())

	r.GET("/ping", controller.ResponseSuccess)

	r.POST("/signup", controller.SignUpHandler)

	r.POST("/login", controller.LoginHandler)

	r.NoRoute(controller.ResponseNotFound)

	return r
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)

		}
		// 处理请求
		c.Next()

	}

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
