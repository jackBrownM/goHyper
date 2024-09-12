package admin_middle

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"goHyper/core/consts"
	"goHyper/internal/ent"
	"goHyper/internal/svc/base"
	"goHyper/libs/errLib"
	"goHyper/libs/jwtLib"
	"strconv"
)

func AdminAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// 获取json
		jwtString := ctx.Cookies(consts.AdminTokenName)
		// 解析json
		adminJwt, err := jwtLib.DecodeAdminJwt(base.GlobalConfig.Admin.JwtSignKey, base.GlobalConfig.Admin.JwtAesKey, jwtString)
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
	idStr := ctx.Get("AdminId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	return id
}

func SetAdminId(ctx *fiber.Ctx, id int) {
	ctx.Set("AdminId", strconv.Itoa(id))
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
