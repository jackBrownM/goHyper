package logic

import (
	req_admin "goHyper/internal/controller/admin/req"
	rsp_admin "goHyper/internal/controller/admin/rsp"
	"goHyper/internal/dao"
	"goHyper/internal/ent"
	"goHyper/libs/errLib"
	"goHyper/libs/resLib"
	"goHyper/svc/base"
	"strings"
)

type Role struct {
	perm   *dao.Perm
	admin  *dao.Admin
	role   *dao.Role
	config *base.Config
}

func NewRole(perm *dao.Perm, admin *dao.Admin, role *dao.Role, config *base.Config) *Role {
	return &Role{
		perm:   perm,
		admin:  admin,
		role:   role,
		config: config,
	}
}

func (l *Role) All() ([]rsp_admin.SystemAuthRoleSimpleRsp, error) {
	return l.role.All()
}

func (l *Role) List(pageReq req_admin.PageReq) (*rsp_admin.PageRsp, error) {
	return l.role.List(pageReq)
}

func (l *Role) Detail(roleId int) (*rsp_admin.SystemAuthRoleRsp, error) {
	roleRsp, err := l.role.Detail(roleId)
	// roleRsp.Member = l.admin.GetMemberCnt(roleId)
	// roleRsp.Menus, err = l.perm.SelectMenuIdsByRoleId(roleId)
	if err != nil {
		return nil, err
	}
	return roleRsp, nil
}

func (l *Role) Create(addReq req_admin.SystemAuthRoleAddReq) error {
	var role *ent.SystemAuthRole
	role, err := l.role.GetByName(strings.Trim(addReq.Name, " "))
	if err != nil {
		return nil
	}
	if role != nil {
		return errLib.RoleNameExist
	}
	resLib.Copy(&role, addReq)
	role.Name = strings.Trim(addReq.Name, " ")
	// ===============================
	// 数据创建
	// ===============================
	err = l.role.Create(role)
	if err != nil {
		return err
	}
	// ===============================
	// 后置操作
	// ===============================
	err = l.perm.BatchSaveByMenuIds(role.Id, addReq.MenuIds)
	if err != nil {
		return err
	}
	return nil
}

func (l *Role) Update(editReq req_admin.SystemAuthRoleEditReq) error {
	// ===============================
	// 前置判断
	// ===============================
	var role *ent.SystemAuthRole
	role, err := l.role.GetByName(strings.Trim(editReq.Name, " "))
	if err != nil {
		return nil
	}
	if role == nil {
		return errLib.RoleNotExist
	}
	isNameExit := l.role.IsNameExit(editReq.ID, editReq.Name)
	if isNameExit {
		return errLib.RoleNameExist
	}
	// ===============================
	// 数据处理
	// ===============================
	role.Name = strings.Trim(editReq.Name, " ")
	// ===============================
	// 数据更新
	// ===============================
	err = l.role.Update(role)
	// ===============================
	// 后置操作
	// ===============================
	err = l.perm.BatchDeleteByRoleId(editReq.ID)
	if err != nil {
		return err
	}
	err = l.perm.BatchSaveByMenuIds(editReq.ID, editReq.MenuIds)
	if err != nil {
		return err
	}
	return nil
}

func (l *Role) Delete(roleId int) error {
	// ===============================
	// 前置判断
	// ===============================
	role, err := l.role.GetById(roleId)
	if err != nil {
		return err
	}
	if role == nil {
		return errLib.RoleNotExist
	}
	isUsed := l.role.IsUsed(roleId)
	if isUsed {
		return errLib.RoleUsed
	}
	// ===============================
	// 数据处理
	// ===============================
	err = l.role.Delete(roleId)
	if err != nil {
		return err
	}
	// ===============================
	// 后置处理
	// ===============================
	if err = l.perm.BatchDeleteByRoleId(roleId); err != nil {
		return err
	}
	return nil

}
