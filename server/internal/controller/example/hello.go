package example

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/internal/service/example"
)

var helloControllerInstance *helloController

type helloController struct{}

func init() {
	helloControllerInstance = &helloController{}
}

func HelloController() *helloController {
	return helloControllerInstance
}

func (h *helloController) SayHello(c *fiber.Ctx) error {
	str := example.HelloService().HelloHay("hello world")
	return c.SendString(str)
}
