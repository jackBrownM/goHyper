package errLib

import "fmt"

type Err struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

func (e Err) Error() string {
	return fmt.Sprintf("errlib:[%d]%s", e.Code, e.Message)
}

func NewErr(code uint32, message string) *Err {
	return &Err{Code: code, Message: message}
}

// GetCode 获取错误代码
func (e Err) GetCode() uint32 {
	return e.Code
}

// GetMsg 获取错误信息
func (e Err) GetMsg() string {
	return e.Message
}

const NotFoundCode = 404
const ServerErrorCode = 500

var (
	NotFound         = NewErr(NotFoundCode, "404 Not found")
	ServerError      = NewErr(ServerErrorCode, "500 Server error")
	Error            = NewErr(100001, "Error")
	Unknown          = NewErr(100002, "Unknown errLib")
	AccountNotExist  = NewErr(100003, "Account not exist")
	AccountLocked    = NewErr(100004, "Account locked")
	AccountDisabled  = NewErr(100005, "Account disabled")
	PasswordError    = NewErr(100006, "Password errLib")
	TokenError       = NewErr(100007, "Token errLib")
	TokenExpired     = NewErr(100008, "Token expired")
	TokenInvalid     = NewErr(100009, "Token invalid")
	IpLocked         = NewErr(200001, "Ip locked")
	IpNotAllow       = NewErr(200002, "Ip not allow")
	PermissionDenied = NewErr(200003, "Permission denied")
	ParamError       = NewErr(200004, "Param errLib")
	AccountExist     = NewErr(200005, "Account is exist")
)
