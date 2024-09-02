package base

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Mysql = gorm.DB

func NewMysql(config *Config, sysLogger *zap.Logger) (*Mysql, error) {
	sysLogger.Info("Gorm初始化...")
	conn := mysql.Open(config.MySQL.Conn)
	gormCfg := &gorm.Config{}
	if config.Svc.IsProd() {
		gormCfg.Logger = logger.Default.LogMode(logger.Error)
	} else {
		gormCfg.Logger = logger.Default.LogMode(logger.Info)
	}
	mysqlDB, err := gorm.Open(conn, gormCfg)
	if err != nil {
		return nil, errors.Wrap(err, "gorm初始化失败")
	}
	sqlDB, err := mysqlDB.DB()
	if err != nil {
		return nil, errors.Wrap(err, "gorm初始化失败")
	}
	sqlDB.SetMaxIdleConns(64)
	sqlDB.SetMaxOpenConns(64)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	return mysqlDB, nil
}
