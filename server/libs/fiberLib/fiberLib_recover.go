package fiberLib

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"goHyper/libs/errLib"
)

func Recover(logger *zap.Logger, hideDetail bool) func(c *fiber.Ctx) (err error) {
	logger = logger.WithOptions(zap.AddCallerSkip(2))
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = errLib.ServerError
				if !hideDetail {
					logger.Sugar().Errorf("%v", r)
				}
			}
		}()
		return c.Next()
	}

}
