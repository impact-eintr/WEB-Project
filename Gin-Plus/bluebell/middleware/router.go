package middleware

import (
	"bluebell/controller"
	"bluebell/pkg/jwt"
	"net/http"
	"strings"

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

// 路由设置
/*
# Cookie-Session认证模式
- 客户端使用用户名，密码进行认证
- 服务端验证用户名、密码正确后生成并存储Session,将SessionID通过Cookie返回客户端
- 客户端访问需要认证的接口时在Cookie中携带SessionID
- 服务端通过SEssionID查找Session并进行鉴权，返回给客户端需要的数据

# Token认证模式
- 客户端使用用户名、密码进行认证
- 服务端验证用户名、密码正确后生成Token后返回客户端
- 客户端保存Token，访问需要认证的接口是URL参数或者HTTP Header中加入Token
- 服务端通过解码Token进行鉴权，返回给客户端需要的数据

*/

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

// 翻译中间件
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

// JWT认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNullAuth)
			c.Abort()
			return

		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidAuth)
			c.Abort()
			return

		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return

		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("userId", mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息

	}

}
