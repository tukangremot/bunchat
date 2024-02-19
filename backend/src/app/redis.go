package app

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func InitRedis(ctx context.Context) (*redis.Client, error) {
	var redisdb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_HOST"),
		Password: viper.GetString("REDIS_PASS"),
		DB:       viper.GetInt("REDIS_DB"),
	})

	err := redisdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return redisdb, nil
}
