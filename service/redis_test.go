package service

import (
	"testing"
)

func TestGetRedis(t *testing.T) {
	config := RedisConfig{
		Addr:     "127.0.0.1:6379",
		Password: "",
	}
	client, err := GetRedis(config)
	if err != nil {
		t.Error(err)
	}
	client.Close()
}
