package dao

import (
	"goHyper/core/svc/base"
	"goHyper/internal/ent"
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
