package ctr_admin

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/core/middleware/admin_ctx"
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
func (c *System) Detail(ctx *fiber.Ctx) error {
	adminId := ctx.QueryInt("id")
	detail, err := c.system.Detail(adminId)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx, detail)
}

func (c *System) Self(ctx *fiber.Ctx) error {
	myId := admin_ctx.GetAdminId(ctx)
	self, err := c.system.Self(myId)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx, self)
}

func (c *System) List(ctx *fiber.Ctx) error {
	var page req_admin.PageReq
	var listReq req_admin.SystemAuthAdminListReq
	if err := httpLib.CheckDTO(ctx, &page); err != nil {
		return err
	}
	if err := httpLib.CheckDTO(ctx, &listReq); err != nil {
		return err
	}
	list, err := c.system.List(page, listReq)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx, list)
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

func (c *System) UpInfo(ctx *fiber.Ctx) error {
	var updateInfo req_admin.SystemAuthAdminUpdateReq
	if err := httpLib.CheckDTO(ctx, &updateInfo); err != nil {
		return err
	}
	adminId := admin_ctx.GetAdminId(ctx)
	err := c.system.UpInfo(adminId, updateInfo)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}

func (c *System) Delete(ctx *fiber.Ctx) error {
	adminId := ctx.QueryInt("id")
	myId := admin_ctx.GetAdminId(ctx)
	err := c.system.Delete(myId, adminId)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}

func (c *System) Disable(ctx *fiber.Ctx) error {
	adminId := ctx.QueryInt("id")
	myId := admin_ctx.GetAdminId(ctx)
	err := c.system.Disable(myId, adminId)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}
