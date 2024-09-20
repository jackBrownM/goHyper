package httpLib

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/libs/errLib"
)

func Success(ctx *fiber.Ctx, data interface{}) error {
	var response map[string]interface{}
	if data == nil {
		response = map[string]interface{}{
			"code": 200,
			"msg":  "成功",
			"data": nil,
		}
	} else {
		response = map[string]interface{}{
			"code": 200,
			"msg":  "成功",
			"data": data,
		}
	}

	return ctx.JSON(response)
}

func Fail(ctx *fiber.Ctx, err *errLib.Err) error {
	response := map[string]interface{}{
		"code": err.Code,
		"msg":  err.Message,
		"data": nil,
	}
	return ctx.JSON(response)
}

// CheckDTO post请求参数保存结构体
func CheckDTO(ctx *fiber.Ctx, req interface{}) error {
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	return nil
}

// CheckGetDTO get请求参数保存结构体
func CheckGetDTO(ctx *fiber.Ctx, req interface{}) error {
	// 绑定查询参数到结构体
	if err := ctx.QueryParser(req); err != nil {
		return err
	}
	return nil
}
