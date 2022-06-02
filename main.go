package main

import (
	"fmt"
	"server/global"
	"server/initialize"
)

func main() {
	// 初始化配置文件
	initialize.Viper()
	// 初始化日志库
	global.Zap = initialize.Zap()
	// 初始化数据库
	global.DB = initialize.Gorm()
	// 注册路由
	routers := initialize.Routers()
	if err := routers.Run(":" + global.CONFIG.Server.Port); err != nil {
		panic(fmt.Errorf("server start failed: %v", err))
	}
}
