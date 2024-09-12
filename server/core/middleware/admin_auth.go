package middleware

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/core/consts"
	"goHyper/core/middleware/admin_ctx"
	"goHyper/internal/svc/base"
	"goHyper/libs/errLib"
	"goHyper/libs/jwtLib"
)

func AdminAuthN() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var jwtString string
		cookieValue := ctx.Cookies(consts.AdminTokenName)
		jwtString = cookieValue
		adminJwt, err := jwtLib.DecodeAdminJwt(base.GlobalConfig.Admin.JwtSignKey, base.GlobalConfig.Admin.JwtAesKey, jwtString)
		if err == nil {
			claims := adminJwt.Claims()
			if claims.Id == 0 {
				return errLib.AdminIdNotFound
			}
			admin_ctx.SetAdminId(ctx, claims.Id)
			return ctx.Next()

		} else {
			return errLib.JwtError
		}
		return ctx.Next()
	}
}
