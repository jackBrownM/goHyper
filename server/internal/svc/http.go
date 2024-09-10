package svc

import (
	"github.com/gofiber/fiber/v2"
	base2 "goHyper/internal/svc/base"
	"goHyper/libs/fiberLib"
)

func NewHttpServ(cfg *base2.Config, logger *base2.Logger) (*fiberLib.HttpServ, error) {
	logger.Info("HttpServ初始化...")
	fiberConfig := fiber.Config{}
	if !cfg.Svc.IsDev() {
		fiberConfig.ProxyHeader = fiber.HeaderXForwardedFor
		fiberConfig.EnableIPValidation = true
	}
	httpServ := fiberLib.NewHttpServ(fiberLib.HttpServProps{
		SvcName: cfg.Svc.Name,
		IsProd:  cfg.Svc.IsProd(),
		Host:    cfg.Svc.Host,
		Port:    cfg.Svc.Port,
		Logger:  logger,
		Config:  fiberConfig,
	})
	return httpServ, nil
}
