package ent

type SystemAuthAdmin struct {
	Id            int    `gorm:"column:id"`              // 主键
	DeptId        int    `gorm:"column:dept_id"`         // 部门ID
	PostId        int    `gorm:"column:post_id"`         // 岗位ID
	Username      string `gorm:"column:username"`        // 用户账号
	Nickname      string `gorm:"column:nickname"`        // 用户昵称
	Password      string `gorm:"column:password"`        // 用户密码
	Avatar        string `gorm:"column:avatar"`          // 用户头像
	Role          string `gorm:"column:role"`            // 角色主键
	Salt          string `gorm:"column:salt"`            // 加密盐巴
	Sort          int    `gorm:"column:sort"`            // 排序编号
	IsMultipoint  int    `gorm:"column:is_multipoint"`   // 多端登录: 0=否, 1=是
	IsDisable     int    `gorm:"column:is_disable"`      // 是否禁用: 0=否, 1=是
	IsDelete      int    `gorm:"column:is_delete"`       // 是否删除: 0=否, 1=是
	LastLoginIp   string `gorm:"column:last_login_ip"`   // 最后登录IP
	LastLoginTime int    `gorm:"column:last_login_time"` // 最后登录
	CreateTime    int    `gorm:"column:create_time"`     // 创建时间
	UpdateTime    int    `gorm:"column:update_time"`     // 更新时间
	DeleteTime    int    `gorm:"column:delete_time"`     // 删除时间
}

func (SystemAuthAdmin) TableName() string {
	return "system_auth_admin"
}
