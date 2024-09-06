package dao

import (
	"goHyper/core/svc/base"
	rsp_admin "goHyper/internal/controller/admin/rsp"
	"goHyper/internal/ent"
	"goHyper/libs/errLib"
	"goHyper/libs/resLib"
)

type Role struct {
	db *base.Mysql
}

func NewRole(db *base.Mysql) *Role {
	return &Role{db: db}
}

func (d *Role) Detail(id int) (*rsp_admin.SystemAuthRoleRsp, error) {
	var role *ent.SystemAuthRole
	var rsp = &rsp_admin.SystemAuthRoleRsp{}
	err := d.db.Model(&ent.SystemAuthRole{}).Where("id = ?", id).Limit(1).First(role).Error
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, errLib.NotFound.Prefix("role id")
	}
	resLib.Copy(rsp, role)
	return rsp, nil
}

func (d *Perm) Get(roleId int) (role ent.SystemAuthRole, err error) {
	err = d.db.Model(&ent.SystemAuthRole{}).Where("id = ? AND is_disable = ?", roleId, 0).Limit(1).First(&role).Error
	return
}
