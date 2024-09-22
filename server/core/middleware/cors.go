package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Cors CORS（跨域资源共享）中间件
func Cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",                              // 替换为实际需要的域名
		AllowHeaders: "*",                              // 替换为实际需要的头部信息
		AllowMethods: "OPTIONS, GET, POST,DELETE, PUT", // 移除重复项
		MaxAge:       3600,                             // 根据实际需求调整
	})
}
