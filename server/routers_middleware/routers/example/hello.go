package routers

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/internal/controller/example"
)

var helloRouterInstance *helloRouter

type helloRouter struct{}

func init() {
	helloRouterInstance = &helloRouter{}
}

func HelloRouter() *helloRouter {
	return helloRouterInstance
}

func (h *helloRouter) InitHello(Router fiber.Router) {
	// 设置分组路由
	mgApi := Router.Group("hello")
	// 注册局部中间件
	// 引入controller层
	ctrl := example.HelloController()
	{
		mgApi.Get("hay", ctrl.SayHello).Name("hello:hay")
	}
}
