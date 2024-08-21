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
func (b *RedisBackend) HasParticipant(ctx context.Context, participantName string) (bool, error) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	res, err := b.client.HGetAll(ctx, fmt.Sprintf("participants:%s", participantName)).Result()
	if err != nil {
		return false, err
	}

	return res == nil, nil
}

// error is returned if participant doesn't exist
func (b *RedisBackend) RegisterParticipant(ctx context.Context, participant *apitypes.Participant) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	_, err := b.client.HSet(ctx, fmt.Sprintf("participants:%s", participant.Username), map[string]interface{}{
		"unsername": participant.Username,
		"password":  participant.Password,
		"email":     participant.EmailAddress,
		"join_time": participant.JoinTime,
	}).Result()
	if err != nil {
		return err
	}

	return nil
}

// this should probably return an error instead if the authorization fails
func (b *RedisBackend) AuthorizeParticipant(context.Context, *apitypes.Participant) (bool, error) {
	return false, nil
}

// experimental (RegesterChannel should return an error if channel is already registered)
func (b *RedisBackend) HasChannel(context.Context, string) (bool, error) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	return false, nil
}

// error is returned if channel doesn't exist
func (b *RedisBackend) RegisterChannel(ctx context.Context, channel *apitypes.Channel) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	_, err := b.client.HSet(ctx, fmt.Sprintf("channels:%s", channel.Name), map[string]interface{}{
		"name":          channel.Name,
		"domain":        channel.Domain,
		"creator":       channel.Creator,
		"creation_time": channel.CreationTime,
	}).Result()
	if err != nil {
		return err
	}

	return nil
}

func (b *RedisBackend) DeleteChannel(ctx context.Context, channelName string) (bool, error) {
	return false, nil
}

func (b *RedisBackend) GetGeneralChatHistory(ctx context.Context) (*apitypes.ChatHistory, error) {
	return &apitypes.ChatHistory{}, nil
}

func (b *RedisBackend) GetChannelHistory(ctx context.Context, channelName string) (*apitypes.ChatHistory, error) {
	return &apitypes.ChatHistory{}, nil
}

func (b *RedisBackend) GetChannelList(ctx context.Context) (*apitypes.ChannelList, error) {
	return &apitypes.ChannelList{}, nil
}

func (b *RedisBackend) GetParticipantList(ctx context.Context) (*apitypes.ParticipanList, error) {
	return &apitypes.ParticipanList{}, nil
}

func (b *RedisBackend) StoreMessage(message *apitypes.ChatMessage) error {
	return nil
}
