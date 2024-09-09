package ent

type SystemAuthMenu struct {
	Id         int         `gorm:"column:id"`          // 主键
	Pid        int         `gorm:"column:pid"`         // 上级菜单
	MenuType   interface{} `gorm:"column:menu_type"`   // 权限类型: M=目录，C=菜单，A=按钮
	MenuName   string      `gorm:"column:menu_name"`   // 菜单名称
	MenuIcon   string      `gorm:"column:menu_icon"`   // 菜单图标
	MenuSort   int         `gorm:"column:menu_sort"`   // 菜单排序
	Perms      string      `gorm:"column:perms"`       // 权限标识
	Paths      string      `gorm:"column:paths"`       // 路由地址
	Component  string      `gorm:"column:component"`   // 前端组件
	Selected   string      `gorm:"column:selected"`    // 选中路径
	Params     string      `gorm:"column:params"`      // 路由参数
	IsCache    int         `gorm:"column:is_cache"`    // 是否缓存: 0=否, 1=是
	IsShow     int         `gorm:"column:is_show"`     // 是否显示: 0=否, 1=是
	IsDisable  int         `gorm:"column:is_disable"`  // 是否禁用: 0=否, 1=是
	CreateTime int         `gorm:"column:create_time"` // 创建时间
	UpdateTime int         `gorm:"column:update_time"` // 更新时间
}

func (SystemAuthMenu) TableName() string {
	return "system_auth_menu"
}
