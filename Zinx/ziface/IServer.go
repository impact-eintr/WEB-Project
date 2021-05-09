package ziface

// 定义一个服务器接口
type IServer interface {
	// 启动
	Start()
	// 停止
	Stop()
	// 运行
	Run()
	// 路由功能
	AddRouter(msgId uint32, router IRouter)
}
