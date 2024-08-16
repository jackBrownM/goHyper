package initialize

import (
	"github.com/gofiber/fiber/v2"
	routers "goHyper/router/routers/example"
)

func RouterInit() *fiber.App {
	// 创建路由
	Routers := fiber.New()
	// 注册基础中间件（前后端功能实现）
	// jwt
	// 鉴权

	// 路由分组
	// *=====================================
	// * example路由组
	// * 例子接口
	// ======================================
	exampleRouter := routers.HelloRouter()
	example := Routers.Group("/example")
	{
		exampleRouter.InitHello(example)
	}

	// *=====================================
	// * api路由组
	// * 用户端接口
	// ======================================

	// *=====================================
	// * admin路由组
	// * 后台端接口
	// ======================================

	return Routers
}
