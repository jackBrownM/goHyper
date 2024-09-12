package base

import (
	"errors"
	"go.uber.org/zap"
)

type Logger = zap.Logger

func NewLogger(cfg *Config) (*Logger, error) {
	var logger *zap.Logger
	var err error
	if cfg.Svc.IsProd() {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, err
	}
	if logger == nil {
		return nil, errors.New("logger base failure")
	}
	if cfg.Svc.IsProd() {
		logger.Info("生产环境 日志初始化成功...")
	} else {
		logger.Info("开发环境 日志初始化成功...")
	}
	return logger, nil
}
