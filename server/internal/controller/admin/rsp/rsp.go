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
	Member     int    `json:"member" structs:"member"`         // 成员数量
	Sort       int    `json:"sort" structs:"sort"`             // 角色排序
	IsDisable  int    `json:"isDisable" structs:"isDisable"`   // 是否禁用: [0=否, 1=是]
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
	IsMultipoint  int    `json:"isMultipoint" structs:"isMultipoint"`   // 多端登录: [0=否, 1=是]
	IsDisable     int    `json:"isDisable" structs:"isDisable"`         // 是否禁用: [0=否, 1=是]
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
	IsMultipoint  int    `json:"isMultipoint" structs:"isMultipoint"`   // 多端登录: [0=否, 1=是]
	IsDisable     int    `json:"isDisable" structs:"isDisable"`         // 是否禁用: [0=否, 1=是]
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

// SystemAuthMenuRsp 系统菜单返回信息
type SystemAuthMenuRsp struct {
	ID         int                 `json:"id" structs:"id"`                       // 主键
	Pid        int                 `json:"pid" structs:"pid"`                     // 上级菜单
	MenuType   string              `json:"menuType" structs:"menuType"`           // 权限类型: [M=目录, C=菜单, A=按钮]
	MenuName   string              `json:"menuName" structs:"menuName"`           // 菜单名称
	MenuIcon   string              `json:"menuIcon" structs:"menuIcon"`           // 菜单图标
	MenuSort   int                 `json:"menuSort" structs:"menuSort"`           // 菜单排序
	Perms      string              `json:"perms" structs:"perms"`                 // 权限标识
	Paths      string              `json:"paths" structs:"paths"`                 // 路由地址
	Component  string              `json:"component" structs:"component"`         // 前端组件
	Selected   string              `json:"selected" structs:"selected"`           // 选中路径
	Params     string              `json:"params" structs:"params"`               // 路由参数
	IsCache    int                 `json:"isCache" structs:"isCache"`             // 是否缓存: [0=否, 1=是]
	IsShow     int                 `json:"isShow" structs:"isShow"`               // 是否显示: [0=否, 1=是]
	IsDisable  int                 `json:"isDisable" structs:"isDisable"`         // 是否禁用: [0=否, 1=是]
	CreateTime int                 `json:"createTime" structs:"createTime"`       // 创建时间
	UpdateTime int                 `json:"updateTime" structs:"updateTime"`       // 更新时间
	Children   []SystemAuthMenuRsp `json:"children,omitempty" structs:"children"` // 子集
}
