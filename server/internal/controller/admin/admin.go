package admin

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/internal/controller/admin/req"
	"goHyper/internal/logic"
	"goHyper/libs/httpLib"
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
	var req req.SystemLoginReq
	if err := httpLib.CheckDTO(ctx, &req); err != nil {
		return err
	}
	// 获取ip地址
	ip := ctx.IP() // 使用 IP() 方法获取 IP 地址
	token, err := c.admin.Login(ctx, req.UserName, req.PassWord, ip)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx, token)
}

func (c *Admin) Logout(ctx *fiber.Ctx) error {
	c.admin.Logout(ctx)
	return resLib.Ok(ctx)
}
