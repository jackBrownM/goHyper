package dao

import (
	"github.com/pkg/errors"
	"goHyper/core/svc/base"
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
