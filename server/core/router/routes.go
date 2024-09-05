package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"goHyper/core/svc/base"
	route_admin "goHyper/internal/api/admin"
)

type Route struct {
	cfg        *base.Config
	logger     *base.Logger
	adminRoute *route_admin.Admin
}

func NewRoute(cfg *base.Config, logger *base.Logger, adminRoute *route_admin.Admin,
) (*Route, error) {
	logger.Info("路由初始化...")
	return &Route{
		cfg:        cfg,
		logger:     logger,
		adminRoute: adminRoute,
	}, nil
}

func (r *Route) Register(app *fiber.App) {
	r.logger.Info("注册路由")
	// ========================
	// admin路由组
	// ========================
	adminGroup := app.Group("/admin")
	adminGroup.Use(healthcheck.New())
	// 路由模块
	r.adminRoute.Register(adminGroup)
	// ========================
	// 打印所有路由
	// ========================
	r.logger.Info("路由注册完成")
	routes := app.GetRoutes(true)
	for _, route := range routes {
		if route.Method != "HEAD" {
			fmt.Printf("%-6s%-50s%s\n", route.Method, route.Path, route.Name)
		}
	}
}
