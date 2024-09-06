package ctr_admin

import (
	"github.com/gofiber/fiber/v2"
	req_admin "goHyper/internal/controller/admin/req"
	"goHyper/internal/logic"
	"goHyper/libs/httpLib"
	"goHyper/libs/resLib"
)

type System struct {
	system *logic.System
}

func NewSystem(system *logic.System) *System {
	return &System{
		system: system,
	}
}

func (c *System) Login(ctx *fiber.Ctx) error {
	var req req_admin.SystemLoginReq
	if err := httpLib.CheckDTO(ctx, &req); err != nil {
		return err
	}
	// 获取ip地址
	ip := ctx.IP() // 使用 IP() 方法获取 IP 地址
	token, err := c.system.Login(ctx, req.UserName, req.PassWord, ip)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx, token)
}

func (c *System) Logout(ctx *fiber.Ctx) error {
	c.system.Logout(ctx)
	return resLib.Ok(ctx)
}

func (c *System) Create(ctx *fiber.Ctx) error {
	var addReq req_admin.SystemAuthAdminAddReq
	if err := httpLib.CheckDTO(ctx, &addReq); err != nil {
		return err
	}
	err := c.system.Create(addReq)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}

func (c *System) Update(ctx *fiber.Ctx) error {
	var editReq req_admin.SystemAuthAdminEditReq
	if err := httpLib.CheckDTO(ctx, &editReq); err != nil {
		return err
	}
	err := c.system.Update(editReq)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}
