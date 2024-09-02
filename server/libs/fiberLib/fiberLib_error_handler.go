package fiberLib

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"goHyper/libs/errLib"
	"goHyper/libs/resLib"
)

const (
	code500 = fiber.StatusInternalServerError
	code400 = fiber.StatusBadRequest
)

func ErrorHandler(logger *zap.Logger, hideDetail bool) func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return c.Status(fiberErr.Code).SendString(fiberErr.Message)
		}
		// Retrieve the custom status code if it's a *fiber.Error
		var libErr *errLib.Err
		if errors.As(err, &libErr) {
			return c.JSON(resLib.Rsp{
				Code: libErr.GetCode(),
				Msg:  libErr.GetMsg(),
				Data: nil,
			})
		}
		if hideDetail {
			// If the error is hidden, print detail in console
			logger.Error("ErrorHandler: unhandled errors")
			fmt.Printf("%+v", errors.WithStack(err))
			return c.Status(code500).JSON(resLib.FromErr(errLib.ServerError))
		} else {
			logger.Error("ErrorHandler: unhandled errors, please view detail in browser")
			// If the error is not hidden, print detail in browser
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			return c.Status(code500).SendString(fmt.Sprintf("%+v", errors.WithStack(err)))
		}
	}

}
