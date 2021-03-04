package routers

import (
	"OSS/apiServer/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//建立子路由处理实例
	temp := v1.NewTemp()
	objects := v1.NewObjects()
	version := v1.NewVersion()
	locate := v1.NewLocate()

	apiv1 := router.Group("api/v1")
	{
		//基础存储服务
		apiv1.PUT("/object/:file", objects.Put)
		apiv1.POST("/object/:file", objects.Post)
		apiv1.GET("/object/:file", objects.Get)
		apiv1.DELETE("/object/:file", objects.Delete)

		//断点续传
		apiv1.PUT("/temp/*file", temp.Put)
		apiv1.HEAD("/temp/*file", temp.Head)

		//版本控制
		apiv1.GET("/versions/:file", version.Get)

		//文件定位服务
		apiv1.GET("/local/:file", locate.Get)
	}

	return router

}
