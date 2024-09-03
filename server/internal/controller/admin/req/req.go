package req

// SystemLoginReq 系统登录参数
type SystemLoginReq struct {
	UserName string `json:"username" binding:"required,min=2,max=20"` // 账号
	PassWord string `json:"password" binding:"required,min=6,max=32"` // 密码
}
