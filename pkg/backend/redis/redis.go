package redis

import (
	"context"
	"fmt"
	"sync"

	"github.com/isnastish/chat-backend/pkg/apitypes"
	"github.com/redis/go-redis/v9"
)

type RedisBackend struct {
	client *redis.Client
	mutex  sync.RWMutex
}

func NewRedisBackend(config *apitypes.RedisConfig) (*RedisBackend, error) {
	// Redis commands docs: https://redis.io/docs/latest/commands/?alpha=x
	opts := redis.Options{
		Addr:     config.Endpoint,
		Username: config.Username,
		Password: config.Password,
	}

	client := redis.NewClient(&opts)
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("Couldn't connect to redis %v", err)
	}

	return &RedisBackend{client: client}, nil
}

// experimental (RegisterParticipant should return an error if participant is already registered)
func (b *RedisBackend) HasParticipant(ctx context.Context, participantName string) bool {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	return false
}

// error is returned if participant doesn't exist
func (b *RedisBackend) RegisterParticipant(context.Context, *apitypes.Participant) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return nil
}

// this should probably return an error instead if the authorization fails
func (b *RedisBackend) AuthorizeParticipant(context.Context, *apitypes.Participant) bool {
	return false
}

// experimental (RegesterChannel should return an error if channel is already registered)
func (b *RedisBackend) HasChannel(context.Context, string) bool {
	return false
}

// error is returned if channel doesn't exist
func (b *RedisBackend) RegisterChannel(context.Context, *apitypes.Channel) error {
	return nil
}

func (b *RedisBackend) DeleteChannel(context.Context, string) bool {
	return false
}

func (b *RedisBackend) GetGeneralChatHistory(context.Context) *apitypes.ChatHistory {
	return &apitypes.ChatHistory{}
}

func (b *RedisBackend) GetChannelHistory(context.Context, string) *apitypes.ChatHistory {
	return &apitypes.ChatHistory{}
}

func (b *RedisBackend) GetChannelList(context.Context) *apitypes.ChannelList {
	return &apitypes.ChannelList{}
}

func (b *RedisBackend) GetParticipantList(context.Context) *apitypes.ParticipanList {
	return &apitypes.ParticipanList{}
}

func (b *RedisBackend) StoreMessage(*apitypes.ChatMessage) {

}
