package ctr_admin

import (
	"github.com/gofiber/fiber/v2"
	req_admin "goHyper/internal/controller/admin/req"
	"goHyper/internal/logic"
	"goHyper/libs/httpLib"
	"goHyper/libs/resLib"
)

type Role struct {
	role *logic.Role
}

func NewRole(role *logic.Role) *Role {
	return &Role{
		role: role,
	}
}

func (c *Role) All(ctx *fiber.Ctx) error {
	all, err := c.role.All()
	if err != nil {
		return err
	}
	return resLib.Success(ctx, all)
}

func (c *Role) List(ctx *fiber.Ctx) error {
	var pageReq req_admin.PageReq
	err := httpLib.CheckDTO(ctx, &pageReq)
	if err != nil {
		return err
	}
	list, err := c.role.List(pageReq)
	if err != nil {
		return err
	}
	return resLib.Success(ctx, list)
}

func (c *Role) Detail(ctx *fiber.Ctx) error {
	roleId := ctx.QueryInt("id")
	detail, err := c.role.Detail(roleId)
	if err != nil {
		return err
	}
	return resLib.Success(ctx, detail)
}

func (c *Role) Create(ctx *fiber.Ctx) error {
	var addReq req_admin.SystemAuthRoleAddReq
	if err := httpLib.CheckDTO(ctx, &addReq); err != nil {
		return err
	}
	err := c.role.Create(addReq)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}

func (c *Role) Update(ctx *fiber.Ctx) error {
	var editReq req_admin.SystemAuthRoleEditReq
	if err := httpLib.CheckDTO(ctx, &editReq); err != nil {
		return err
	}
	err := c.role.Update(editReq)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)

}

func (c *Role) Delete(ctx *fiber.Ctx) error {
	roleId := ctx.QueryInt("id")
	err := c.role.Delete(roleId)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}
