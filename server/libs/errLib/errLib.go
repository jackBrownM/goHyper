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
const (
	Success = iota
	ErrorCode
	UnknownCode
	AccountNotExistCode
	AccountLockedCode
	AccountDisabledCode
	PasswordErrorCode
	TokenErrorCode
	TokenExpiredCode
	TokenInvalidCode
	IpLockedCode
	IpNotAllowCode
	PermissionDeniedCode
	ParamErrorCode
)

var (
	NotFound         = NewErr(NotFoundCode, "404 Not found")
	ServerError      = NewErr(ServerErrorCode, "500 Server error")
	Error            = NewErr(ErrorCode, "Error")
	Unknown          = NewErr(UnknownCode, "Unknown errLib")
	AccountNotExist  = NewErr(AccountNotExistCode, "Account not exist")
	AccountLocked    = NewErr(AccountLockedCode, "Account locked")
	AccountDisabled  = NewErr(AccountDisabledCode, "Account disabled")
	PasswordError    = NewErr(PasswordErrorCode, "Password errLib")
	TokenError       = NewErr(TokenErrorCode, "Token errLib")
	TokenExpired     = NewErr(TokenExpiredCode, "Token expired")
	TokenInvalid     = NewErr(TokenInvalidCode, "Token invalid")
	IpLocked         = NewErr(IpLockedCode, "Ip locked")
	IpNotAllow       = NewErr(IpNotAllowCode, "Ip not allow")
	PermissionDenied = NewErr(PermissionDeniedCode, "Permission denied")
	ParamError       = NewErr(ParamErrorCode, "Param errLib")
)
