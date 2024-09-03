package base

import (
	"go.uber.org/zap"
	"goHyper/libs/cfgLib"
	"goHyper/libs/mysqlLib"
	"goHyper/libs/redisLib"
	"goHyper/libs/svcLib"
)

type Jwt struct {
	JwtSignKey string
	JwtAesKey  string
}

type Config struct {
	Svc   svcLib.Config
	MySQL mysqlLib.Config
	Redis redisLib.Config
	Admin Jwt
}

func NewConfig() (*Config, error) {
	logger, _ := zap.NewDevelopment()
	logger.Info("Config初始化...")
	cfg, err := cfgLib.Load[Config](logger)
	return cfg, err
}
