package example

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/internal/logic/example"
)

type helloController struct {
}

func HelloController() *helloController {
	return &helloController{}
}

// SayHello 控制器层只做代码的接收验证
func (h *helloController) SayHello(c *fiber.Ctx) error {
	str := example.HelloService().HelloHay("hello world")
	return c.SendString(str)
}
