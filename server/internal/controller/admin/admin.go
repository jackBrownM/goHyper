package admin

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/internal/logic"
	"goHyper/libs/resLib"
)

type Admin struct {
	admin *logic.Admin
}

func NewAdmin(admin *logic.Admin) *Admin {
	return &Admin{
		admin: admin,
	}
}

func (c *Admin) Login(ctx *fiber.Ctx) error {
	str := c.admin.Login()
	return resLib.Ok(ctx, str)
}
