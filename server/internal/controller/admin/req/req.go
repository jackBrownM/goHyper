package req_admin

// SystemLoginReq 系统登录参数
type SystemLoginReq struct {
	UserName string `json:"username" binding:"required,min=2,max=20"` // 账号
	PassWord string `json:"password" binding:"required,min=6,max=32"` // 密码
}

// SystemAuthAdminAddReq 管理员新增参数
type SystemAuthAdminAddReq struct {
	DeptId       int    `form:"deptId" binding:"required,gt=0"`           // 部门ID
	PostId       int    `form:"postId" binding:"required,gt=0"`           // 岗位ID
	Username     string `form:"username" binding:"required,min=2,max=20"` // 账号
	Nickname     string `form:"nickname" binding:"required,min=2,max=30"` // 昵称
	Password     string `form:"password" binding:"required,min=6,max=20"` // 密码
	Avatar       string `form:"avatar" binding:"required"`                // 头像
	Role         int    `form:"role" binding:"gte=0"`                     // 角色
	Sort         int    `form:"sort" binding:"gte=0"`                     // 排序
	IsDisable    int8   `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
	IsMultipoint int8   `form:"isMultipoint" binding:"oneof=0 1"`         // 多端登录: [0=否, 1=是]
}

// SystemAuthAdminEditReq 管理员编辑参数
type SystemAuthAdminEditReq struct {
	ID           int    `form:"id" binding:"required,gt=0"`               // 主键
	DeptId       int    `form:"deptId" binding:"required,gt=0"`           // 部门ID
	PostId       int    `form:"postId" binding:"required,gt=0"`           // 岗位ID
	Username     string `form:"username" binding:"required,min=2,max=20"` // 账号
	Nickname     string `form:"nickname" binding:"required,min=2,max=30"` // 昵称
	Password     string `form:"password"`                                 // 密码
	Avatar       string `form:"avatar"`                                   // 头像
	Role         int    `form:"role" binding:"gte=0"`                     // 角色
	Sort         int    `form:"sort" binding:"gte=0"`                     // 排序
	IsDisable    int8   `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
	IsMultipoint int8   `form:"isMultipoint" binding:"oneof=0 1"`         // 多端登录: [0=否, 1=是]
}

// SystemAuthAdminUpdateReq 管理员更新参数
type SystemAuthAdminUpdateReq struct {
	Nickname     string `form:"nickname" binding:"required,min=2,max=30"`     // 昵称
	Avatar       string `form:"avatar"`                                       // 头像
	Password     string `form:"password" binding:"required"`                  // 密码
	CurrPassword string `form:"currPassword" binding:"required,min=6,max=32"` // 密码
}

// PageReq 分页请求参数
type PageReq struct {
	PageNo   int `form:"pageNo,default=1" validate:"omitempty,gte=1"`          // 页码
	PageSize int `form:"pageSize,default=20" validate:"omitempty,gt=0,lte=60"` // 每页大小
}

// SystemAuthAdminListReq 管理员列表参数
type SystemAuthAdminListReq struct {
	Username string `form:"username"`        // 账号
	Nickname string `form:"nickname"`        // 昵称
	Role     int    `form:"role,default=-1"` // 角色ID
}
