package rsp_admin

// SystemLoginRsp 系统登录返回信息
type SystemLoginRsp struct {
	Token string `json:"token"`
}

// SystemAuthRoleRsp 系统角色返回信息
type SystemAuthRoleRsp struct {
	ID         uint   `json:"id" structs:"id"`                 // 主键
	Name       string `json:"name" structs:"name"`             // 角色名称
	Remark     string `json:"remark" structs:"remark"`         // 角色备注
	Menus      []int  `json:"menus" structs:"menus"`           // 关联菜单
	Member     int64  `json:"member" structs:"member"`         // 成员数量
	Sort       uint16 `json:"sort" structs:"sort"`             // 角色排序
	IsDisable  uint8  `json:"isDisable" structs:"isDisable"`   // 是否禁用: [0=否, 1=是]
	CreateTime int    `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime int    `json:"updateTime" structs:"updateTime"` // 更新时间
}
