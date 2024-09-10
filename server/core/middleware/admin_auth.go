package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"goHyper/core/consts"
	"goHyper/core/middleware/admin_ctx"
	"goHyper/internal/svc/base"
	"goHyper/libs/errLib"
	"goHyper/libs/jwtLib"
)

func AdminAuthN(config *base.Jwt) gin.HandlerFunc {
	return func(gc *fiber.Ctx) {
		var jwtString string
		if cookieValue, err := gc.Cookie(consts.AdminTokenName); err == nil {
			jwtString = cookieValue
		} else if headValue := gc.GetHeader(consts.AdminTokenName); headValue != "" {
			jwtString = headValue
		}
		adminJwt, err := jwtLib.DecodeAdminJwt(config.JwtSignKey, config.JwtAesKey, jwtString)
		if err == nil {
			claims := adminJwt.Claims()
			if claims.Id == 0 {
				panic(errLib.AdminIdNotFound)
			}
			admin_ctx.SetAdminId(gc, claims.Id)
			gc.Next()
			return
		} else {
			panic(errLib.JwtError)
		}
		panic(errLib.NotLogin)
		return
	}
}
