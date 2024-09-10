package dao

import (
	"goHyper/internal/ent"
	"goHyper/internal/svc/base"
	"goHyper/libs/utilLib"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type Perm struct {
	db *base.Mysql
}

func NewPerm(db *base.Mysql) *Perm {
	return &Perm{db: db}
}

// SelectMenuIdsByRoleId 根据角色ID获取菜单ID
func (d *Perm) SelectMenuIdsByRoleId(roleId int) (menuIds []int, e error) {
	var role ent.SystemAuthRole
	err := d.db.Where("id = ? AND is_disable = ?", roleId, 0).Limit(1).First(&role).Error
	if err != nil {
		return nil, err
	}
	var perms []ent.SystemAuthPerm
	err = d.db.Where("role_id = ?", role.Id).Find(&perms).Error
	if err != nil {
		return nil, err
	}
	for _, perm := range perms {
		menuIds = append(menuIds, perm.MenuId)
	}
	return
}

func (d *Perm) GetListByRoleId(roleId int) (perms []ent.SystemAuthPerm, err error) {
	err = d.db.Model(&ent.SystemAuthPerm{}).Where("role_id = ?", roleId).Find(&perms).Error
	return
}

func (d *Perm) BatchDeleteByRoleId(roleId int) (err error) {
	err = d.db.Delete(&ent.SystemAuthPerm{}, "role_id = ?", roleId).Error
	return
}

func (d *Perm) BatchSaveByMenuIds(roleId int, menuIds string) (err error) {
	if menuIds == "" {
		return
	}

	err = d.db.Transaction(func(tx *gorm.DB) error {
		var perms []ent.SystemAuthPerm
		for _, menuIdStr := range strings.Split(menuIds, ",") {
			menuId, _ := strconv.ParseInt(menuIdStr, 10, 32)
			perms = append(perms, ent.SystemAuthPerm{Id: utilLib.MakeUuid(), RoleId: roleId, MenuId: int(menuId)})
		}
		err = tx.Create(&perms).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return
}
