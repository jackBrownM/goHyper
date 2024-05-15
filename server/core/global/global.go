package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	HyperDB  *gorm.DB
	HyperLog *zap.Logger
)
