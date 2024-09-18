package ctr_admin

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/core/middleware/admin_middle"
	req_admin "goHyper/internal/controller/admin/req"
	"goHyper/internal/logic"
	"goHyper/libs/httpLib"
	"goHyper/libs/resLib"
)

type Menu struct {
	menu *logic.Menu
}

func NewMenu(menu *logic.Menu) *Menu {
	return &Menu{
		menu: menu,
	}
}

func (c *Menu) Route(ctx *fiber.Ctx) error {
	myId := admin_middle.GetAdminId(ctx)
	list, err := c.menu.SelectMenuByRoleId(myId)
	if err != nil {
		return err
	}
	// return ctx.JSON(map[string]interface{}{
	//	"code": 200,
	//	"msg":  "成功",
	//	"data": list,
	// })

	return resLib.Success(ctx, list)
}

func (c *Menu) List(ctx *fiber.Ctx) error {
	list, err := c.menu.List()
	if err != nil {
		return err
	}
	return resLib.Success(ctx, list)
}

func (c *Menu) Detail(ctx *fiber.Ctx) error {
	id := ctx.QueryInt("id")
	detail, err := c.menu.Detail(id)
	if err != nil {
		return err
	}
	return resLib.Success(ctx, detail)
}

func (c *Menu) Create(ctx *fiber.Ctx) error {
	var req req_admin.SystemAuthMenuAddReq
	err := httpLib.CheckDTO(ctx, &req)
	if err != nil {
		return err
	}
	err = c.menu.Create(req)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}

func (c *Menu) Update(ctx *fiber.Ctx) error {
	var editReq req_admin.SystemAuthMenuEditReq
	err := c.menu.Update(editReq)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}

func (c *Menu) Delete(ctx *fiber.Ctx) error {
	menuId := ctx.QueryInt("id")
	err := c.menu.Delete(menuId)
	if err != nil {
		return err
	}
	return resLib.Ok(ctx)
}
