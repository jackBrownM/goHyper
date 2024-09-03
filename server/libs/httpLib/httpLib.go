package httpLib

import (
	"github.com/gofiber/fiber/v2"
)

func CheckDTO(ctx *fiber.Ctx, req interface{}) error {
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	return nil
}
