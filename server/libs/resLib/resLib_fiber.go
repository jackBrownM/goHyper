package resLib

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/libs/errLib"
)

func Ok(ctx *fiber.Ctx, data ...any) error {
	if len(data) == 0 {
		return ctx.JSON(Rsp{
			Code: 0,
			Msg:  "success",
			Data: nil,
		})
	}
	return ctx.JSON(Rsp{
		Code: 0,
		Msg:  "success",
		Data: data[0],
	})
}

func Error(ctx *fiber.Ctx, err *errLib.Err) error {
	return ctx.JSON(Rsp{
		Code: err.Code,
		Msg:  err.Message,
		Data: struct{}{},
	})
}

func FromErr(err *errLib.Err) *Rsp {
	return &Rsp{
		Code: err.GetCode(),
		Msg:  err.GetMsg(),
		Data: nil,
	}
}
