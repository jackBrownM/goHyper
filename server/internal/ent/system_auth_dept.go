package ent

type SystemAuthDept struct {
	Id         int    `gorm:"column:id"`          // 主键
	Pid        int    `gorm:"column:pid"`         // 上级主键
	Name       string `gorm:"column:name"`        // 部门名称
	Duty       string `gorm:"column:duty"`        // 负责人名
	Mobile     string `gorm:"column:mobile"`      // 联系电话
	Sort       int    `gorm:"column:sort"`        // 排序编号
	IsStop     int    `gorm:"column:is_stop"`     // 是否禁用: 0=否, 1=是
	IsDelete   int    `gorm:"column:is_delete"`   // 是否删除: 0=否, 1=是
	CreateTime int    `gorm:"column:create_time"` // 创建时间
	UpdateTime int    `gorm:"column:update_time"` // 更新时间
	DeleteTime int    `gorm:"column:delete_time"` // 删除时间
}

func (SystemAuthDept) TableName() string {
	return "system_auth_dept"
}
