package dao

import (
	req_admin "goHyper/internal/controller/admin/req"
	rsp_admin "goHyper/internal/controller/admin/rsp"
	"goHyper/internal/ent"
	"goHyper/libs/errLib"
	"goHyper/svc/base"
	"time"
)

type Admin struct {
	db *base.Mysql
}

func NewAdmin(db *base.Mysql) *Admin {
	return &Admin{db: db}
}

func (d *Admin) Create(sysAdmin ent.SystemAuthAdmin) (err error) {
	err = d.db.Create(&sysAdmin).Error
	return
}

func (d *Admin) Update(adminMap ent.SystemAuthAdmin) (err error) {
	err = d.db.Model(ent.SystemAuthAdmin{}).Where("id = ?", adminMap.Id).Updates(adminMap).Error
	return
}

// GetByUserName 根据账号查找管理员
func (d *Admin) GetByUserName(userName string) (admin ent.SystemAuthAdmin, err error) {
	err = d.db.Model(ent.SystemAuthAdmin{}).Where("username = ? and is_delete <> 1", userName).Limit(1).First(&admin).Error
	if err != nil {
		err = errLib.AccountNotExist
	}
	return
}

// LoginUpdate 登录更新
func (d *Admin) LoginUpdate(adminId int, ip string) (err error) {
	err = d.db.Model(ent.SystemAuthAdmin{}).Where("id = ?", adminId).Updates(ent.SystemAuthAdmin{LastLoginIp: ip, LastLoginTime: int(time.Now().Unix())}).Error
	return
}

// IsExitAdmin 通过用户名判断是否存在admin
func (d *Admin) IsExitAdmin(userName string, nickName string) bool {
	var count int64
	d.db.Model(ent.SystemAuthAdmin{}).Where("(username = ? or nickname = ? ) and is_delete <> 1", userName, nickName).Count(&count)
	return count > 0
}

// GetMemberCnt 根据角色ID获取成员数量
func (d *Admin) GetMemberCnt(roleId int) int64 {
	var count int64
	d.db.Model(&ent.SystemAuthAdmin{}).Where(
		"role = ? AND is_delete = ?", roleId, 0).Count(&count)
	return count
}

// GetById 根据id查找admin
func (d *Admin) GetById(id int) (*ent.SystemAuthAdmin, error) {
	var admin ent.SystemAuthAdmin
	err := d.db.Model(&ent.SystemAuthAdmin{}).Where("id = ? AND is_delete = 0", id).Limit(1).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, err
}

// Delete 删除管理员
func (d *Admin) Delete(id int) (err error) {
	err = d.db.Model(ent.SystemAuthAdmin{}).Where("id = ?", id).Updates(ent.SystemAuthAdmin{IsDelete: 1, DeleteTime: int(time.Now().Unix())}).Error
	return
}

func (d *Admin) Disable(id, isDisable int) (err error) {
	err = d.db.Model(ent.SystemAuthAdmin{}).Where("id = ?", id).Updates(ent.SystemAuthAdmin{IsDisable: isDisable, UpdateTime: int(time.Now().Unix())}).Error
	return
}

func (d *Admin) List(page req_admin.PageReq) (*rsp_admin.PageRsp, error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 构建查询模型
	adminModel := d.db.Table(ent.TableNameSystemAuthAdmin+" AS admin").
		Joins("LEFT JOIN "+ent.TableNameSystemAuthRole+" role ON admin.role = role.id").
		Where("admin.is_delete = ?", 0)
	// 总数
	var count int64
	err := adminModel.Count(&count).Error
	if err != nil {
		return nil, err
	}
	// 数据
	var adminRsp []rsp_admin.SystemAuthAdminRsp
	err = adminModel.Limit(limit).Offset(offset).
		Order("admin.id DESC, admin.sort DESC").
		Select("admin.*, role.name as role").
		Find(&adminRsp).Error
	if err != nil {
		return nil, err
	}
	// 处理
	for i := 0; i < len(adminRsp); i++ {
		if adminRsp[i].ID == 1 {
			adminRsp[i].Role = "系统管理员"
		}
	}
	return &rsp_admin.PageRsp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    adminRsp,
	}, nil
}
