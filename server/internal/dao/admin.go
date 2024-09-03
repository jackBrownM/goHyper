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

func (d *Admin) GetUser() string {
	return "user:xxxxxx"
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
