package req_admin

// SystemLoginReq 系统登录参数
type SystemLoginReq struct {
	UserName string `json:"username" binding:"required,min=2,max=20"` // 账号
	PassWord string `json:"password" binding:"required,min=6,max=32"` // 密码
}

// SystemAuthAdminAddReq 管理员新增参数
type SystemAuthAdminAddReq struct {
	Username     string `form:"username" binding:"required,min=2,max=20"` // 账号
	Nickname     string `form:"nickname" binding:"required,min=2,max=30"` // 昵称
	Password     string `form:"password" binding:"required,min=6,max=20"` // 密码
	Avatar       string `form:"avatar" binding:"required"`                // 头像
	Role         int    `form:"role" binding:"gte=0"`                     // 角色
	Sort         int    `form:"sort" binding:"gte=0"`                     // 排序
	IsDisable    int    `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
	IsMultipoint int    `form:"isMultipoint" binding:"oneof=0 1"`         // 多端登录: [0=否, 1=是]
}

// SystemAuthAdminEditReq 管理员编辑参数
type SystemAuthAdminEditReq struct {
	ID           int    `form:"id" binding:"required,gt=0"`               // 主键
	Username     string `form:"username" binding:"required,min=2,max=20"` // 账号
	Nickname     string `form:"nickname" binding:"required,min=2,max=30"` // 昵称
	Password     string `form:"password"`                                 // 密码
	Role         int    `form:"role" binding:"gte=0"`                     // 角色
	Sort         int    `form:"sort" binding:"gte=0"`                     // 排序
	IsDisable    int    `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
	IsMultipoint int8   `form:"isMultipoint" binding:"oneof=0 1"`         // 多端登录: [0=否, 1=是]
}

// SystemAuthAdminUpdateReq 管理员更新参数
type SystemAuthAdminUpdateReq struct {
	Nickname     string `form:"nickname" binding:"required,min=2,max=30"`     // 昵称
	Avatar       string `form:"avatar"`                                       // 头像
	Password     string `form:"password" binding:"required"`                  // 密码
	CurrPassword string `form:"currPassword" binding:"required,min=6,max=32"` // 密码
}

// SystemAuthAdminDelReq 管理员删除参数
type SystemAuthAdminDelReq struct {
	ID int `form:"id" binding:"required,gt=0"` // 主键
}

// SystemAuthRoleDelReq 角色删除参数
type SystemAuthRoleDelReq struct {
	ID int `form:"id" binding:"required,gt=0"` // 主键
}

// PageReq 分页请求参数
type PageReq struct {
	PageNo   int `query:"pageNo,default=1" validate:"omitempty,gte=1"`          // 页码
	PageSize int `query:"pageSize,default=20" validate:"omitempty,gt=0,lte=60"` // 每页大小
}

// SystemAuthAdminListReq 管理员列表参数
type SystemAuthAdminListReq struct {
	Username string `form:"username"`        // 账号
	Nickname string `form:"nickname"`        // 昵称
	Role     int    `form:"role,default=-1"` // 角色ID
}

// SystemAuthRoleAddReq 新增角色参数
type SystemAuthRoleAddReq struct {
	Name      string `form:"name" binding:"required,min=1,max=30"` // 角色名称
	Sort      int    `form:"sort" binding:"gte=0"`                 // 角色排序
	IsDisable int    `form:"isDisable" binding:"oneof=0 1"`        // 是否禁用: [0=否, 1=是]
	Remark    string `form:"remark" binding:"max=200"`             // 角色备注
	MenuIds   string `form:"menuIds"`                              // 关联菜单
}

// SystemAuthRoleEditReq 编辑角色参数
type SystemAuthRoleEditReq struct {
	ID        int    `form:"id" binding:"required,gt=0"`           // 主键
	Name      string `form:"name" binding:"required,min=1,max=30"` // 角色名称
	Sort      int    `form:"sort" binding:"gte=0"`                 // 角色排序
	IsDisable int8   `form:"isDisable" binding:"oneof=0 1"`        // 是否禁用: [0=否, 1=是]
	Remark    string `form:"remark" binding:"max=200"`             // 角色备注
	MenuIds   string `form:"menuIds"`                              // 关联菜单
}

// SystemAuthMenuAddReq 新增菜单参数
type SystemAuthMenuAddReq struct {
	Pid       uint   `form:"pid" binding:"gte=0"`                      // 上级菜单
	MenuType  string `form:"menuType" binding:"oneof=M C A"`           // 权限类型: [M=目录, C=菜单, A=按钮]
	MenuName  string `form:"menuName" binding:"required,min=1,max=30"` // 菜单名称
	MenuIcon  string `form:"menuIcon" binding:"max=100"`               // 菜单图标
	MenuSort  int    `form:"menuSort" binding:"gte=0"`                 // 菜单排序
	Perms     string `form:"perms" binding:"max=100"`                  // 权限标识
	Paths     string `form:"paths" binding:"max=200"`                  // 路由地址
	Component string `form:"component" binding:"max=200"`              // 前端组件
	Selected  string `form:"selected" binding:"max=200"`               // 选中路径
	Params    string `form:"params" binding:"max=200"`                 // 路由参数
	IsCache   uint8  `form:"isCache" binding:"oneof=0 1"`              // 是否缓存: [0=否, 1=是]
	IsShow    uint8  `form:"isShow" binding:"oneof=0 1"`               // 是否显示: [0=否, 1=是]
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
}

// SystemAuthMenuEditReq 编辑菜单参数
type SystemAuthMenuEditReq struct {
	ID        int    `form:"id" binding:"required,gt=0"`               // 主键
	Pid       int    `form:"pid" binding:"gte=0"`                      // 上级菜单
	MenuType  string `form:"menuType" binding:"oneof=M C A"`           // 权限类型: [M=目录, C=菜单, A=按钮]
	MenuName  string `form:"menuName" binding:"required,min=1,max=30"` // 菜单名称
	MenuIcon  string `form:"menuIcon" binding:"max=100"`               // 菜单图标
	MenuSort  int    `form:"menuSort" binding:"gte=0"`                 // 菜单排序
	Perms     string `form:"perms" binding:"max=100"`                  // 权限标识
	Paths     string `form:"paths" binding:"max=200"`                  // 路由地址
	Component string `form:"component" binding:"max=200"`              // 前端组件
	Selected  string `form:"selected" binding:"max=200"`               // 选中路径
	Params    string `form:"params" binding:"max=200"`                 // 路由参数
	IsCache   int    `form:"isCache" binding:"oneof=0 1"`              // 是否缓存: [0=否, 1=是]
	IsShow    int    `form:"isShow" binding:"oneof=0 1"`               // 是否显示: [0=否, 1=是]
	IsDisable int    `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
}

// SystemAuthMenuDelReq 删除菜单参数
type SystemAuthMenuDelReq struct {
	ID int `form:"id" binding:"required,gt=0"` // 主键
}
