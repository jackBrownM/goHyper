package route_admin

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/internal/controller/admin"
)

type Admin struct {
	system *ctr_admin.Admin
	role   *ctr_admin.Role
}

func NewAdmin(system *ctr_admin.Admin, role *ctr_admin.Role) *Admin {
	return &Admin{system: system, role: role}
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
		adminGroup.Get("/admin/detail", r.system.Detail).Name("管理员详情")
		adminGroup.Get("/admin/list", r.system.List).Name("管理员列表")
		adminGroup.Get("/admin/self", r.system.Self).Name("管理员自己的信息")
		adminGroup.Post("/admin/add", r.system.Create).Name("管理员创建")
		adminGroup.Post("/admin/edit", r.system.Update).Name("管理员更新")
		adminGroup.Post("/admin/upInfo", r.system.UpInfo).Name("管理员信息更新")
		adminGroup.Get("/admin/delete", r.system.Delete).Name("管理员删除")
		adminGroup.Get("/admin/disable", r.system.Disable).Name("管理员状态切换")
	}
	// 角色
	{
		adminGroup.Get("/role/all", r.role.All).Name("获取所有角色")
		adminGroup.Get("/role/list", r.role.List).Name("角色列表")
		adminGroup.Get("/role/detail", r.role.Detail).Name("角色详情")
		adminGroup.Post("/role/add", r.role.Create).Name("角色创建")
		adminGroup.Post("/role/edit", r.role.Update).Name("角色更新")
		adminGroup.Get("/role/delete", r.role.Delete).Name("角色删除")
	}

}
