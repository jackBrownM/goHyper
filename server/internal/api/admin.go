package api

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/internal/controller/admin"
)

type Admin struct {
	admin *admin.Admin
}

func NewAdmin(admin *admin.Admin) *Admin {
	return &Admin{admin: admin}
}

func (r *Admin) Register(root fiber.Router) {
	adminGroup := root.Group("/system")
	{
		adminGroup.Post("/login", r.admin.Login).Name("登录系统")
		adminGroup.Post("/logout", r.admin.Logout).Name("退出系统")
	}
}
