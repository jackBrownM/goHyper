package logic

import (
	"goHyper/core/svc/base"
	req_admin "goHyper/internal/controller/admin/req"
	rsp_admin "goHyper/internal/controller/admin/rsp"
	"goHyper/internal/dao"
	"goHyper/internal/ent"
	"goHyper/libs/errLib"
	"goHyper/libs/resLib"
	"goHyper/libs/utilLib"
)

type Menu struct {
	menu   *dao.Menu
	config *base.Config
}

func NewMenu(menu *dao.Menu, config *base.Config) *Menu {
	return &Menu{
		menu:   menu,
		config: config,
	}
}

func (l *Menu) SelectMenuByRoleId(adminId int) (mapList []interface{}, err error) {
	menuIds, err := l.menu.SelectMenuIdsByRoleId(adminId)
	if err != nil {
		return nil, err
	}

	if len(menuIds) == 0 {
		menuIds = []int{0}
	}

	menus, err := l.menu.GetChain()
	if err != nil {
		return nil, err
	}
	var menuRsps []rsp_admin.SystemAuthMenuRsp
	resLib.Copy(&menuRsps, menus)
	mapList = utilLib.ArrayUtil.ListToTree(
		utilLib.ConvertUtil.StructsToMaps(menuRsps), "id", "pid", "children")
	return nil, nil
}

func (l *Menu) List() (list []interface{}, err error) {
	menus, err := l.menu.List()
	var menuRsps []rsp_admin.SystemAuthMenuRsp
	resLib.Copy(&menuRsps, menus)
	list = utilLib.ArrayUtil.ListToTree(
		utilLib.ConvertUtil.StructsToMaps(menuRsps), "id", "pid", "children")
	return
}

func (l *Menu) Detail(id int) (detail *rsp_admin.SystemAuthMenuRsp, err error) {
	menu, err := l.menu.GetById(id)
	if err != nil {
		return nil, err
	}
	resLib.Copy(detail, menu)
	return
}

func (l *Menu) Create(addReq req_admin.SystemAuthMenuAddReq) (err error) {
	var menu ent.SystemAuthMenu
	resLib.Copy(&menu, addReq)
	err = l.menu.Create(menu)
	return
}

func (l *Menu) Update(editReq req_admin.SystemAuthMenuEditReq) (err error) {
	menu, err := l.menu.GetById(editReq.ID)
	if err != nil {
		return
	}
	if menu == nil {
		return errLib.NotFound.Prefix("菜单不存在")
	}
	resLib.Copy(menu, editReq)
	err = l.menu.Update(menu)
	if err != nil {
		return
	}
	return
}

func (l *Menu) Delete(menuId int) (err error) {
	menu, err := l.menu.GetById(menuId)
	if err != nil {
		return
	}
	if menu == nil {
		return errLib.NotFound.Prefix("菜单不存在")
	}
	menuPid, err := l.menu.GetByPid(menuId)
	if err != nil {
		return
	}
	if menuPid != nil {
		return errLib.CannotDeleteMenu.Prefix("该菜单下存在子菜单，无法删除")
	}
	err = l.menu.Delete(menuId)
	if err != nil {
		return
	}
	return
}
