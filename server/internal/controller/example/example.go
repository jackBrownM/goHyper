package example

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/internal/logic"
	"goHyper/libs/resLib"
)

type Example struct {
	example *logic.Example
}

func NewExample(example *logic.Example) *Example {
	return &Example{
		example: example,
	}
}

func (c *Example) Example(ctx *fiber.Ctx) error {
	str := c.example.Example()
	return resLib.Ok(ctx, str)
}
