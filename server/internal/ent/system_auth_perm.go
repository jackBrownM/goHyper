package ent

type SystemAuthPerm struct {
	Id     string `gorm:"column:id"`      // 主键
	RoleId int    `gorm:"column:role_id"` // 角色ID
	MenuId int    `gorm:"column:menu_id"` // 菜单ID
}

func (SystemAuthPerm) TableName() string {
	return "system_auth_perm"
}
