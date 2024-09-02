package api

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/internal/controller/example"
)

type Example struct {
	example *example.Example
}

func NewExample(example *example.Example) *Example {
	return &Example{example: example}
}

func (r *Example) Register(root fiber.Router, prefix string) {
	exampleGroup := root.Group(prefix)
	{
		// 登录系统
		exampleGroup.Get("/example", r.example.Example).Name("读取系统信息")
	}
}
