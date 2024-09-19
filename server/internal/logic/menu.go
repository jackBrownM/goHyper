package logic

import (
	req_admin "goHyper/internal/controller/admin/req"
	rsp_admin "goHyper/internal/controller/admin/rsp"
	"goHyper/internal/dao"
	"goHyper/internal/ent"
	"goHyper/libs/resLib"
	"goHyper/libs/utilLib"
	"goHyper/svc/base"
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
	return mapList, nil
}

func (l *Menu) List() (list []interface{}, err error) {
	menus, err := l.menu.List()
	var menuRsps []rsp_admin.SystemAuthMenuRsp
	resLib.Copy(&menuRsps, menus)
	list = utilLib.ArrayUtil.ListToTree(
		utilLib.ConvertUtil.StructsToMaps(menuRsps), "id", "pid", "children")
	return
}

func (l *Menu) Detail(id int) (*rsp_admin.SystemAuthMenuRsp, error) {
	var detail rsp_admin.SystemAuthMenuRsp
	menu, err := l.menu.GetById(id)
	if err != nil {
		return nil, err
	}
	resLib.Copy(&detail, menu)
	return &detail, nil
}

func (l *Menu) Create(addReq req_admin.SystemAuthMenuAddReq) (err error) {
	var menu ent.SystemAuthMenu
	resLib.Copy(&menu, addReq)
	err = l.menu.Create(menu)
	return
}

func (l *Menu) Update(editReq req_admin.SystemAuthMenuEditReq) (err error) {
	//menu, err := l.menu.GetById(editReq.ID)
	//if err != nil {
	//	return
	//}
	//if menu == nil {
	//	return errLib.NotFound.Prefix("菜单不存在")
	//}
	var menu ent.SystemAuthMenu
	resLib.Copy(&menu, editReq)
	err = l.menu.Update(&menu)
	if err != nil {
		return
	}
	return
}

func (l *Menu) Delete(menuId int) (err error) {
	err = l.menu.Delete(menuId)
	if err != nil {
		return
	}
	return
}
