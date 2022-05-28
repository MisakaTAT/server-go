package main

import (
	"server/global"
	"server/initialize"
)

func main() {
	// 初始化日志库
	global.Zap = initialize.Zap()
	// 初始化配置文件
	global.Viper = initialize.Viper()
	// 初始化数据库
	//global.DB = initialize.Gorm()
	// 注册路由
	routers := initialize.Routers()
	if err := routers.Run(":" + global.CONFIG.Server.Port); err != nil {
		global.Zap.Panicf("Server start failed: %v", err)
	}
}
