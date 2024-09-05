package route_admin

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/internal/controller/admin"
)

type Admin struct {
	system *ctr_admin.System
}

func NewAdmin(system *ctr_admin.System) *Admin {
	return &Admin{system: system}
}

func (r *Admin) Register(root fiber.Router) {
	adminGroup := root.Group("/system")
	// 登录
	{
		adminGroup.Post("/login", r.system.Login).Name("登录系统")
		adminGroup.Post("/logout", r.system.Logout).Name("退出系统")
	}
	// 管理员
	{
		adminGroup.Post("/admin/add", r.system.Create)
	}
}
