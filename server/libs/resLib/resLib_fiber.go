package resLib

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
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
		Data: data,
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

func CookieRemove(ctx *fiber.Ctx, cookieName string) {
	ctx.Cookie(&fiber.Cookie{
		Name:     cookieName,
		Value:    "",
		MaxAge:   0,
		Path:     "",
		Domain:   "",
		Secure:   false,
		HTTPOnly: true,
	})
}

func CookieAdd(ctx *fiber.Ctx, cookieName string, cookieValue string, maxAge int) {
	ctx.Cookie(&fiber.Cookie{
		Name:     cookieName,
		Value:    cookieValue,
		MaxAge:   3600 * maxAge,
		Path:     "",
		Domain:   "",
		Secure:   false,
		HTTPOnly: true,
	})
}

// Copy 拷贝结构体
func Copy(toValue interface{}, fromValue interface{}) interface{} {
	if err := copier.Copy(toValue, fromValue); err != nil {
		panic(err)
	}
	return toValue
}
