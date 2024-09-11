package admin_ctx

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"goHyper/internal/ent"
	"strconv"
)

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
