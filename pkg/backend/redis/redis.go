package redis

import (
	"github.com/isnastish/chat-backend/pkg/apitypes"
	"github.com/redis/go-redis/v9"
)

type RedisBackend struct {
	client *redis.Client
}

func NewRedisBackend(condig *apitypes.RedisConfig) (*RedisBackend, error) {
	return &RedisBackend{}, nil
}
