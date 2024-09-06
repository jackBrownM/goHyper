package ent

type SystemAuthRole struct {
	Id         int    `gorm:"column:id"`          // 主键
	Name       string `gorm:"column:name"`        // 角色名称
	Remark     string `gorm:"column:remark"`      // 备注信息
	Sort       int    `gorm:"column:sort"`        // 角色排序
	IsDisable  int    `gorm:"column:is_disable"`  // 是否禁用: 0=否, 1=是
	CreateTime int    `gorm:"column:create_time"` // 创建时间
	UpdateTime int    `gorm:"column:update_time"` // 更新时间
}

func (SystemAuthRole) TableName() string {
	return "system_auth_role"
}
