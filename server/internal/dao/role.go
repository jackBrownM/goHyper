package dao

import (
	"goHyper/core/svc/base"
	req_admin "goHyper/internal/controller/admin/req"
	rsp_admin "goHyper/internal/controller/admin/rsp"
	"goHyper/internal/ent"
	"goHyper/libs/errLib"
	"goHyper/libs/resLib"
	"gorm.io/gorm"
	"strings"
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

func (d *Role) All() (rsp []rsp_admin.SystemAuthRoleSimpleRsp, err error) {
	var roles []ent.SystemAuthRole
	err = d.db.Model(&ent.SystemAuthRole{}).Order("sort desc, id desc").Find(&roles).Error
	if err != nil {
		return
	}
	resLib.Copy(&rsp, roles)
	return
}

func (d *Role) List(page req_admin.PageReq) (*rsp_admin.PageRsp, error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	roleModel := d.db.Model(&ent.SystemAuthRole{})
	var count int64
	err := roleModel.Count(&count).Error
	if err != nil {
		return nil, err
	}
	var roles []ent.SystemAuthRole
	err = roleModel.Limit(limit).Offset(offset).Order("sort desc, id desc").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	var roleResp []rsp_admin.SystemAuthRoleRsp
	resLib.Copy(&roleResp, roles)
	for i := 0; i < len(roleResp); i++ {
		roleResp[i].Menus = []int{}
		roleResp[i].Member = d.getMemberCnt(roleResp[i].ID)
	}
	return &rsp_admin.PageRsp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    roleResp,
	}, nil
}

// getMemberCnt 根据角色ID获取成员数量
func (d *Role) getMemberCnt(roleId int) (count int64) {
	d.db.Model(&ent.SystemAuthAdmin{}).Where(
		"role = ? AND is_delete = ?", roleId, 0).Count(&count)
	return
}

func (d *Role) GetByName(name string) (role *ent.SystemAuthRole, err error) {
	err = d.db.Model(&ent.SystemAuthRole{}).Where("name = ?", name).Limit(1).First(role).Error
	return
}

func (d *Role) GetById(id int) (role *ent.SystemAuthRole, err error) {
	err = d.db.Model(&ent.SystemAuthRole{}).Where("id = ?", id).Limit(1).First(role).Error
	return
}

func (d *Role) Create(role *ent.SystemAuthRole) error {
	return d.db.Create(role).Error
}

func (d *Role) IsNameExit(id int, name string) bool {
	var count int64
	d.db.Model(&ent.SystemAuthRole{}).Where("id != ? AND name = ?", id, strings.Trim(name, " ")).Count(&count)
	return count > 0
}

func (d *Role) Update(role *ent.SystemAuthRole) (err error) {
	err = d.db.Model(&ent.SystemAuthRole{}).Where("id = ?", role.Id).Updates(role).Error
	return
}

func (d *Role) IsUsed(roleId int) bool {
	var count int64
	d.db.Model(&ent.SystemAuthRole{}).Where("role = ? AND is_delete = ?", roleId, 0).Count(&count)
	return count > 0
}

func (d *Role) Delete(id int) (err error) {
	// 事务
	err = d.db.Transaction(func(tx *gorm.DB) error {
		err = tx.Delete(&ent.SystemAuthRole{}, "id = ?", id).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
	return
}
