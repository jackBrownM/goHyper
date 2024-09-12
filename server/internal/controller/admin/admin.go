package ctr_admin

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/core/middleware/admin_ctx"
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
	var req req_admin.SystemLoginReq
	if err := httpLib.CheckDTO(ctx, &req); err != nil {
		return err
	}
	// 获取ip地址
	ip := ctx.IP() // 使用 IP() 方法获取 IP 地址
	token, err := c.admin.Login(ctx, req.UserName, req.PassWord, ip)
	if err != nil {
		return err
	}
	return ctx.JSON(map[string]interface{}{
		"code": 200,
		"msg":  "成功",
		"data": token,
	})
}

func (c *Admin) Logout(ctx *fiber.Ctx) error {
	c.admin.Logout(ctx)
	return resLib.Ok(ctx)
}

// 管理员
func (c *Admin) Detail(ctx *fiber.Ctx) error {
	adminId := ctx.QueryInt("id")
	detail, err := c.admin.Detail(adminId)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx, detail)
}

func (c *Admin) Self(ctx *fiber.Ctx) error {
	myId := admin_middle.GetAdminId(ctx)
	self, err := c.admin.Self(myId)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx, self)
}

func (c *Admin) List(ctx *fiber.Ctx) error {
	var page req_admin.PageReq
	var listReq req_admin.SystemAuthAdminListReq
	if err := httpLib.CheckDTO(ctx, &page); err != nil {
		return err
	}
	if err := httpLib.CheckDTO(ctx, &listReq); err != nil {
		return err
	}
	list, err := c.admin.List(page, listReq)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx, list)
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
	adminId := admin_ctx.GetAdminId(ctx)
	err := c.admin.UpInfo(adminId, updateInfo)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}

func (c *Admin) Delete(ctx *fiber.Ctx) error {
	adminId := ctx.QueryInt("id")
	myId := admin_ctx.GetAdminId(ctx)
	err := c.admin.Delete(myId, adminId)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}

func (c *Admin) Disable(ctx *fiber.Ctx) error {
	adminId := ctx.QueryInt("id")
	myId := admin_ctx.GetAdminId(ctx)
	err := c.admin.Disable(myId, adminId)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}
