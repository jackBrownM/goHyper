package ctr_admin

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/core/middleware/admin_middle"
	req_admin "goHyper/internal/controller/admin/req"
	"goHyper/internal/logic"
	"goHyper/libs/httpLib"
	"goHyper/libs/resLib"
)

type Admin struct {
	admin *logic.Admin
}

func NewSystem(admin *logic.Admin) *Admin {
	return &Admin{
		admin: admin,
	}
}

// Login 登录
func (c *Admin) Login(ctx *fiber.Ctx) error {
	// 结构体校验与获取
	var req req_admin.SystemLoginReq
	if err := httpLib.CheckDTO(ctx, &req); err != nil {
		return err
	}
	// 获取ip地址
	ip := ctx.IP()
	// 登录逻辑
	token, err := c.admin.Login(req.UserName, req.PassWord, ip)
	if err != nil {
		return err
	}
	// 返回消息
	return httpLib.Success(ctx, token)
}

func (c *Admin) Logout(ctx *fiber.Ctx) error {
	c.admin.Logout(ctx)
	return resLib.Ok(ctx)
}

func (c *Admin) Detail(ctx *fiber.Ctx) error {
	adminId := ctx.QueryInt("id")
	detail, err := c.admin.Detail(adminId)
	if err != nil {
		return err
	}
	return resLib.Success(ctx, detail)
}

func (c *Admin) Self(ctx *fiber.Ctx) error {
	myId := admin_middle.GetAdminId(ctx)
	self, err := c.admin.Self(myId)
	if err != nil {
		return err
	}
	return httpLib.Success(ctx, self)
}

func (c *Admin) List(ctx *fiber.Ctx) error {
	var page req_admin.PageReq
	page.PageNo = ctx.QueryInt("pageNo")
	page.PageSize = ctx.QueryInt("pageSize")
	list, err := c.admin.List(page)
	if err != nil {
		return err
	}
	return httpLib.Success(ctx, list)
}

func (c *Admin) Create(ctx *fiber.Ctx) error {
	var addReq req_admin.SystemAuthAdminAddReq
	if err := httpLib.CheckDTO(ctx, &addReq); err != nil {
		return err
	}
	err := c.admin.Create(addReq)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}

func (c *Admin) Update(ctx *fiber.Ctx) error {
	var editReq req_admin.SystemAuthAdminEditReq
	if err := httpLib.CheckDTO(ctx, &editReq); err != nil {
		return err
	}
	err := c.admin.Update(editReq)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}

func (c *Admin) UpInfo(ctx *fiber.Ctx) error {
	var updateInfo req_admin.SystemAuthAdminUpdateReq
	if err := httpLib.CheckDTO(ctx, &updateInfo); err != nil {
		return err
	}
	adminId := admin_middle.GetAdminId(ctx)
	err := c.admin.UpInfo(adminId, updateInfo)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}

func (c *Admin) Delete(ctx *fiber.Ctx) error {
	var delInfo req_admin.SystemAuthAdminDelReq
	if err := httpLib.CheckDTO(ctx, &delInfo); err != nil {
		return err
	}
	myId := admin_middle.GetAdminId(ctx)
	err := c.admin.Delete(myId, delInfo.ID)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}

func (c *Admin) Disable(ctx *fiber.Ctx) error {
	adminId := ctx.QueryInt("id")
	myId := admin_middle.GetAdminId(ctx)
	err := c.admin.Disable(myId, adminId)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}
