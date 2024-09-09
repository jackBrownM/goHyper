package rsp_admin

// SystemLoginRsp 系统登录返回信息
type SystemLoginRsp struct {
	Token string `json:"token"`
}

// SystemAuthRoleRsp 系统角色返回信息
type SystemAuthRoleRsp struct {
	ID         int    `json:"id" structs:"id"`                 // 主键
	Name       string `json:"name" structs:"name"`             // 角色名称
	Remark     string `json:"remark" structs:"remark"`         // 角色备注
	Menus      []int  `json:"menus" structs:"menus"`           // 关联菜单
	Member     int64  `json:"member" structs:"member"`         // 成员数量
	Sort       int16  `json:"sort" structs:"sort"`             // 角色排序
	IsDisable  int8   `json:"isDisable" structs:"isDisable"`   // 是否禁用: [0=否, 1=是]
	CreateTime int    `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime int    `json:"updateTime" structs:"updateTime"` // 更新时间
}

// SystemAuthAdminSelfOneRsp 当前管理员返回部分信息
type SystemAuthAdminSelfOneRsp struct {
	ID            int    `json:"id" structs:"id"`                       // 主键
	Username      string `json:"username" structs:"username"`           // 账号
	Nickname      string `json:"nickname" structs:"nickname"`           // 昵称
	Avatar        string `json:"avatar" structs:"avatar"`               // 头像
	Role          string `json:"role" structs:"role"`                   // 角色
	Dept          string `json:"dept" structs:"dept"`                   // 部门
	IsMultipoint  int8   `json:"isMultipoint" structs:"isMultipoint"`   // 多端登录: [0=否, 1=是]
	IsDisable     int8   `json:"isDisable" structs:"isDisable"`         // 是否禁用: [0=否, 1=是]
	LastLoginIp   string `json:"lastLoginIp" structs:"lastLoginIp"`     // 最后登录IP
	LastLoginTime int    `json:"lastLoginTime" structs:"lastLoginTime"` // 最后登录时间
	CreateTime    int    `json:"createTime" structs:"createTime"`       // 创建时间
	UpdateTime    int    `json:"updateTime" structs:"updateTime"`       // 更新时间
}

// SystemAuthAdminSelfRsp 当前系统管理员返回信息
type SystemAuthAdminSelfRsp struct {
	User        SystemAuthAdminSelfOneRsp `json:"user" structs:"user"`               // 用户信息
	Permissions []string                  `json:"permissions" structs:"permissions"` // 权限集合: [[*]=>所有权限, ['article:add']=>部分权限]
}

// PageRsp 分页返回值
type PageRsp struct {
	Count    int64       `json:"count"`    // 总数
	PageNo   int         `json:"pageNo"`   // 页No
	PageSize int         `json:"pageSize"` // 每页Size
	Lists    interface{} `json:"lists"`    // 数据
}

// SystemAuthAdminRsp 管理员返回信息
type SystemAuthAdminRsp struct {
	ID            int    `json:"id" structs:"id"`                       // 主键
	Username      string `json:"username" structs:"username"`           // 账号
	Nickname      string `json:"nickname" structs:"nickname"`           // 昵称
	Avatar        string `json:"avatar" structs:"avatar"`               // 头像
	Role          string `json:"role" structs:"role"`                   // 角色
	DeptId        uint   `json:"deptId" structs:"deptId"`               // 部门ID
	PostId        uint   `json:"postId" structs:"postId"`               // 岗位ID
	Dept          string `json:"dept" structs:"dept"`                   // 部门
	IsMultipoint  uint8  `json:"isMultipoint" structs:"isMultipoint"`   // 多端登录: [0=否, 1=是]
	IsDisable     uint8  `json:"isDisable" structs:"isDisable"`         // 是否禁用: [0=否, 1=是]
	LastLoginIp   string `json:"lastLoginIp" structs:"lastLoginIp"`     // 最后登录IP
	LastLoginTime int    `json:"lastLoginTime" structs:"lastLoginTime"` // 最后登录时间
	CreateTime    int    `json:"createTime" structs:"createTime"`       // 创建时间
	UpdateTime    int    `json:"updateTime" structs:"updateTime"`       // 更新时间
}

// SystemAuthRoleSimpleRsp 系统角色返回简单信息
type SystemAuthRoleSimpleRsp struct {
	ID         int    `json:"id" structs:"id"`                 // 主键
	Name       string `json:"name" structs:"name"`             // 角色名称
	CreateTime int    `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime int    `json:"updateTime" structs:"updateTime"` // 更新时间
}
