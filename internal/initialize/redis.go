package initialize

import (
	"context"
	"ecom-project/global"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	// Redis initialization
	redisConfig := global.Config.Redis
	addr := fmt.Sprintf("%s:%v", redisConfig.Host, redisConfig.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Protocol: redisConfig.Protocol,
		DB:       redisConfig.DB,
		PoolSize: 15, // default is 10 connections for each CPU
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization failed", zap.Error(err))
		return
	}

	global.Redis = rdb
	redisExample()
}

func redisExample() {
	var rdc = global.Redis
	rdc.Set(ctx, "EcomDemo", "Anhdung123", 0)
	value, err := rdc.Get(ctx, "EcomDemo").Result()
	if err != nil {
		global.Logger.Error((err).Error(), zap.String("redis", "zap"))
	}

	global.Logger.Info((value), zap.String("redis", "zap"))
}
