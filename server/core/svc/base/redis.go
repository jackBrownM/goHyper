package base

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"strconv"
)

type Redis = redis.Client

func NewRedis(config *Config, logger *zap.Logger) (*Redis, error) {
	logger.Info("Redis初始化...")
	var redisDB = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + strconv.Itoa(config.Redis.Port),
		DB:       config.Redis.DBNo,
		Password: config.Redis.Password,
	})
	statusCmd := redisDB.Ping(context.Background())
	if statusCmd.Err() != nil {
		return nil, statusCmd.Err()
	}
	return redisDB, nil
}
