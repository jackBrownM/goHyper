package req_admin

// SystemLoginReq 系统登录参数
type SystemLoginReq struct {
	UserName string `json:"username" binding:"required,min=2,max=20"` // 账号
	PassWord string `json:"password" binding:"required,min=6,max=32"` // 密码
}

// SystemAuthAdminAddReq 管理员新增参数
type SystemAuthAdminAddReq struct {
	DeptId       uint   `form:"deptId" binding:"required,gt=0"`           // 部门ID
	PostId       uint   `form:"postId" binding:"required,gt=0"`           // 岗位ID
	Username     string `form:"username" binding:"required,min=2,max=20"` // 账号
	Nickname     string `form:"nickname" binding:"required,min=2,max=30"` // 昵称
	Password     string `form:"password" binding:"required"`              // 密码
	Avatar       string `form:"avatar" binding:"required"`                // 头像
	Role         uint   `form:"role" binding:"gte=0"`                     // 角色
	Sort         int    `form:"sort" binding:"gte=0"`                     // 排序
	IsDisable    uint8  `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
	IsMultipoint uint8  `form:"isMultipoint" binding:"oneof=0 1"`         // 多端登录: [0=否, 1=是]
}
