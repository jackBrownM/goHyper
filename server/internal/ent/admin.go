package ent

import (
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	ID           int64          `gorm:"column:id;pk;auto"`    // 主键
	No           string         `gorm:"column:no"`            // 编号
	Account      string         `gorm:"column:account"`       // 账号
	Password     string         `gorm:"column:password"`      // 密码
	PasswordSalt string         `gorm:"column:password_salt"` // 密码slat
	Username     string         `gorm:"column:username"`      // 用户名
	Email        string         `gorm:"column:email"`         // 邮箱
	Tel          string         `gorm:"column:tel"`           // 电话
	Role         int8           `gorm:"column:role"`          // 权限
	Enabled      bool           `gorm:"column:enabled"`       // 是否启用
	CreatedAt    time.Time      `gorm:"column:created_at"`    // 创建时间
	CreatedBy    string         `gorm:"column:created_by"`    // 创建者
	LoginAt      time.Time      `gorm:"column:login_at"`      // 登录时间
	LoginIp      string         `gorm:"column:login_ip"`      // 登录ip
	LoginTimes   int32          `gorm:"column:login_times"`   // 登录次数
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`    // 删除时间
}

func (Admin) TableName() string {
	return "admin"
}
