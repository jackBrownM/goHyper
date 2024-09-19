package dao

import (
	"goHyper/internal/ent"
	"goHyper/svc/base"
)

type Menu struct {
	db *base.Mysql
}

func NewMenu(db *base.Mysql) *Menu {
	return &Menu{db: db}
}

func (d *Menu) GetListByRoleId(menuIds []int) (menus []ent.SystemAuthMenu, err error) {
	err = d.db.Model(&ent.SystemAuthMenu{}).Where("id in ? AND is_disable = ? AND menu_type in ?", menuIds, 0, []string{"C", "A"}).Order(
		"menu_sort, id").First(&menus).Error
	return
}

func (d *Menu) SelectMenuIdsByRoleId(roleId int) (menuIds []int, err error) {
	var role ent.SystemAuthRole
	err = d.db.Model(&role).Where("id = ? AND is_disable = ?", roleId, 0).Limit(1).First(&role).Error
	if err != nil {
		return []int{}, err
	}
	var perms []ent.SystemAuthPerm
	err = d.db.Model(&ent.SystemAuthPerm{}).Where("role_id = ?", role.Id).Find(&perms).Error
	if err != nil {
		return []int{}, err
	}
	for _, perm := range perms {
		menuIds = append(menuIds, perm.MenuId)
	}
	return
}

func (d *Menu) GetChain() ([]ent.SystemAuthMenu, error) {
	var menus []ent.SystemAuthMenu
	err := d.db.Table("system_auth_menu").Where("menu_type in ? AND is_disable = ?", []string{"M", "C"}, 0).Order("menu_sort desc, id").Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (d *Menu) List() ([]ent.SystemAuthMenu, error) {
	var menus []ent.SystemAuthMenu
	err := d.db.Model(&ent.SystemAuthMenu{}).Order("menu_sort desc, id").Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (d *Menu) GetById(id int) (*ent.SystemAuthMenu, error) {
	var menu ent.SystemAuthMenu
	err := d.db.Model(&ent.SystemAuthMenu{}).Where("id = ?", id).Limit(1).First(&menu).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (d *Menu) GetByPid(id int) (menu *ent.SystemAuthMenu, err error) {
	err = d.db.Model(&ent.SystemAuthMenu{}).Where("pid = ?", id).Limit(1).First(menu).Error
	return
}

func (d *Menu) Create(menu ent.SystemAuthMenu) (err error) {
	err = d.db.Create(&menu).Error
	return
}

func (d *Menu) Update(menu *ent.SystemAuthMenu) error {
	err := d.db.Model(ent.SystemAuthMenu{}).Where("id = ?", menu.Id).Updates(menu).Error
	return err
}

func (d *Menu) Delete(id int) (err error) {
	err = d.db.Delete(&ent.SystemAuthMenu{}, id).Error
	return
}
