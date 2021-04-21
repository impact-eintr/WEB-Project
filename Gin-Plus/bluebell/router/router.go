package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middleware"
	"bluebell/setting"

	"github.com/gin-gonic/gin"
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
func Setup() *gin.Engine {
	if setting.Conf.AppConfig.Mode == "release" {
		gin.SetMode(gin.ReleaseMode) // 设置为发布模式
	}

	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.Cors())

	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)

	r.GET("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		controller.ResponseSuccess(c, nil)
	})

	r.GET("/home", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		controller.ResponseSuccess(c, nil)
	})

	r.NoRoute(controller.ResponseNotFound)

	return r
}
