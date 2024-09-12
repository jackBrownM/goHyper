package svc

import (
	"github.com/gofiber/fiber/v2"
	"goHyper/libs/fiberLib"
	"goHyper/svc/base"
)

func NewHttpServ(cfg *base.Config, logger *base.Logger) (*fiberLib.HttpServ, error) {
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
