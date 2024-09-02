package zapLib

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(mode string) *zap.Logger {
	// 定义日志文件配置
	logCfg := lumberjack.Logger{
		Filename:   "application.log", // 日志文件路径
		MaxSize:    100,               // 单个日志文件最大大小（以MB为单位）
		MaxBackups: 3,                 // 保留旧文件的最大数量
		MaxAge:     28,                // 旧文件保留天数
		Compress:   true,              // 是否压缩旧文件
	}
	var level zapcore.Level
	if mode == "prod" {
		level = zap.InfoLevel
	} else {
		level = zap.DebugLevel
	}
	// 初始化 zap.Logger
	writeSyncer := zapcore.AddSync(&logCfg)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 使用 JSON 编码器
		writeSyncer,
		level, // 设置最低日志级别
	)
	logger := zap.New(core, zap.AddCaller()) // 添加调用者信息
	defer logger.Sync()
	return logger
}
