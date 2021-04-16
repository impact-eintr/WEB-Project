package main

import (
	"bluebell/logger"
	sf "bluebell/pkg/snowflake"
	"bluebell/setting"
	"fmt"
)

func main() {
	// 1. 加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Printf("init mysql failed, err:#{err}\n")
		return
	}

	// 2. 初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init mysql failed, err:#{err}\n")
		return
	}

	// 3. 初始化MySQL连接
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:#{err}\n")
		return
	}

	// 4. 初始化缓存
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:#{err}\n")
		return
	}

	// 初始化ID生成器
	if err := sf.Init("2021-04-16", setting.Conf.MachineID); err != nil {
		fmt.Printf("init failed,err:%v\n", err)
		return
	}
	id := sf.GenID()
	fmt.Println(id)

	// 5. 注册路由
	//r := router.SetuoRouter()
	//err := r.Run(fmt.Sprintf(":#{setting.AppConfig.Port}"))
	//if err != nil {
	//	fmt.Printf("run server failed, err:#{err}")
	//	return
	//}

	// 6. 启动服务(优雅关机)
}
