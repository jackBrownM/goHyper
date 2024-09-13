package admin_middle

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/libs/errLib"
	"goHyper/libs/jwtLib"
	"goHyper/svc/base"
)

// AdminAuth 后台用户鉴权
func AdminAuth() func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// 尝试从header获取token
		token := ctx.Get("token")
		if len(token) == 0 {
			return errLib.NullValue.Prefix("jwt")
		}
		// 解析json
		adminJwt, err := jwtLib.DecodeAdminJwt(base.GlobalConfig.Admin.JwtSignKey, base.GlobalConfig.Admin.JwtAesKey, token)
		if err != nil {
			return errLib.JwtError
		}
		// 获取结构体
		claims := adminJwt.Claims()
		if claims.Id == 0 {
			return errLib.AdminIdNotFound
		}
		// 向上下文设置adminId
		SetAdminId(ctx, claims.Id)
		return ctx.Next()
	}
}

// GetAdminId 从上下文获取adminId
func GetAdminId(ctx *fiber.Ctx) int {
	id := ctx.Locals("AdminId")
	return id.(int)
}

// SetAdminId 向上下文设置adminId
func SetAdminId(ctx *fiber.Ctx, id int) {
	ctx.Locals("AdminId", id)
}
