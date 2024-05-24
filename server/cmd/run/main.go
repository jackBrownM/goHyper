package main

import (
	"goHyper/core/global"
	initialize2 "goHyper/core/initialize"
)

// 入口文件
func main() {
	// 初始化日志
	global.HyperLog = initialize2.LogInit()
	// 初始化viper
	initialize2.ViperInit()
	// 初始化数据库
	global.HyperDB = initialize2.GormMysql()
	// 初始化路由
	router := initialize2.RouterInit()
	// 启动
	router.Listen(":8081")
}
