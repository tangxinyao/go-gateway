package service

import (
	"github.com/go-redis/redis"
)

func GetRedis(config RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return client, nil
}
