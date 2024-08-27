package main

import (
	"goHyper/core/global"
	initialize "goHyper/core/initialize"
)

// 入口文件
func main() {
	// 初始化日志
	global.HyperLog = initialize.LogInit()
	// 初始化viper
	initialize.ViperInit()
	// 初始化数据库
	// global.HyperDB = initialize.GormMysql()
	// 初始化路由
	router := initialize.RouterInit()
	// 启动
	router.Listen(":8081")
}
