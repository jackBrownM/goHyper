package dao

import (
	"fmt"
	"github.com/pkg/errors"
	"goHyper/core/svc/base"
	req_admin "goHyper/internal/controller/admin/req"
	rsp_admin "goHyper/internal/controller/admin/rsp"
	"goHyper/internal/ent"
	"goHyper/libs/errLib"
	"gorm.io/gorm"
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
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errLib.AccountNotExist
	}
	return
}

// LoginUpdate 登录更新
func (d *Admin) LoginUpdate(ip string) (err error) {
	err = d.db.Model(ent.SystemAuthAdmin{}).Updates(ent.SystemAuthAdmin{LastLoginIp: ip, LastLoginTime: int(time.Now().Unix())}).Error
	return
}

// IsExitAdmin 通过用户名判断是否存在admin
func (d *Admin) IsExitAdmin(userName string, nickName string) bool {
	var count int64
	d.db.Model(ent.SystemAuthAdmin{}).Where("username = ? or nickname = ? and is_delete <> 1", userName, nickName).Count(&count)
	return count > 0
}

// GetMemberCnt 根据角色ID获取成员数量
func (d *Admin) GetMemberCnt(roleId int) (count int64) {
	d.db.Model(&ent.SystemAuthAdmin{}).Where(
		"role = ? AND is_delete = ?", roleId, 0).Count(&count)
	return
}

// GetById 根据id查找admin
func (d *Admin) GetById(id int) (admin *ent.SystemAuthAdmin, err error) {
	err = d.db.Model(admin).Where("id = ? AND is_delete <> 0", id).Limit(1).First(admin).Error
	return admin, err
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

func (d *Admin) List(page req_admin.PageReq, listReq req_admin.SystemAuthAdminListReq) (*rsp_admin.PageRsp, error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	adminTbName := ent.SystemAuthAdmin{}.TableName()
	roleTbName := ent.SystemAuthRole{}.TableName()
	deptTbName := ent.SystemAuthDept{}.TableName()
	adminModel := d.db.Table(adminTbName+" AS admin").Where("admin.is_delete = ?", 0).Joins(
		fmt.Sprintf("LEFT JOIN %s ON admin.role = %s.id", roleTbName, roleTbName)).Joins(
		fmt.Sprintf("LEFT JOIN %s ON admin.dept_id = %s.id", deptTbName, deptTbName)).Select(
		fmt.Sprintf("admin.*, %s.name as dept, %s.name as role", deptTbName, roleTbName))
	// 条件
	if listReq.Username != "" {
		adminModel = adminModel.Where("username like ?", "%"+listReq.Username+"%")
	}
	if listReq.Nickname != "" {
		adminModel = adminModel.Where("nickname like ?", "%"+listReq.Nickname+"%")
	}
	if listReq.Role >= 0 {
		adminModel = adminModel.Where("role = ?", listReq.Role)
	}
	// 总数
	var count int64
	err := adminModel.Count(&count).Error
	if err != nil {
		return nil, err
	}
	// 数据
	var adminRsp []rsp_admin.SystemAuthAdminRsp
	err = adminModel.Limit(limit).Offset(offset).Order("id desc, sort desc").Find(&adminRsp).Error
	if err != nil {
		return nil, err
	}
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
