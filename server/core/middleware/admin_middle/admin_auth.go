package admin_middle

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"goHyper/internal/ent"
	"goHyper/libs/errLib"
	"goHyper/libs/jwtLib"
	"goHyper/svc/base"
)

func AdminAuth() func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// 获取json
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

func GetAdminId(ctx *fiber.Ctx) int {
	id := ctx.Locals("AdminId")
	return id.(int)
}

func SetAdminId(ctx *fiber.Ctx, id int) {
	ctx.Locals("AdminId", id)
}

func GetAdmin(ctx *gin.Context) *ent.SystemAuthAdmin {
	obj, exists := ctx.Get("Admin")
	if exists {
		if admin, ok := obj.(*ent.SystemAuthAdmin); ok {
			return admin
		}
	}
	return nil
}

func SetAdmin(ctx *gin.Context, member *ent.SystemAuthAdmin) {
	ctx.Set("Admin", member)
}
