package cfgLib

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"goHyper/libs/fileLib"
)

func Load[T any](logger *zap.Logger) (*T, error) {
	configPath := ""
	checkPaths := []string{
		"./etc/config.dev.yaml",
		"./etc/config.test.yaml",
		"./etc/config.prod.yaml",
	}
	for _, checkPath := range checkPaths {
		if fileLib.Exists(checkPath) {
			configPath = checkPath
			break
		}
	}
	if configPath == "" {
		return nil, errors.New("Config fileLib not found")
	}
	logger.Info("加载配置文件：" + configPath)
	vip := viper.New()
	vip.SetConfigType("yaml")
	vip.SetConfigFile(configPath)

	if err := vip.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "Viper read errLib")
	}

	c := new(T)
	if err := vip.Unmarshal(c); err != nil {
		return nil, errors.Wrap(err, "Viper unmarshal errLib")
	}
	return c, nil
}
