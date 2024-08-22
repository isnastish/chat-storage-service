package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/isnastish/chat-backend/pkg/apitypes"
	"github.com/isnastish/chat-backend/pkg/utils"
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
		return nil, fmt.Errorf("couldn't connect to redis %v", err)
	}

	return &RedisBackend{client: client}, nil
}

func (b *RedisBackend) HasParticipant(ctx context.Context, participantName string) (bool, error) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	isMember, err := b.client.SIsMember(ctx, "participants:", participantName).Result()
	return isMember, err
}

func (b *RedisBackend) RegisterParticipant(ctx context.Context, participant *apitypes.Participant) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	hashedPassword, _ := utils.SHA256(participant.Password)

	if _, err := b.client.HSet(ctx, participant.Username, map[string]interface{}{
		"unsername": participant.Username,
		"password":  hashedPassword,
		"email":     participant.EmailAddress,
		"join_time": participant.JoinTime,
	}).Result(); err != nil {
		return err
	}
	// Add participant's username to `participants` set
	_, err := b.client.SAdd(ctx, "participants:", participant.Username).Result()
	return err
}

func (b *RedisBackend) AuthorizeParticipant(ctx context.Context, participant *apitypes.Participant) (bool, error) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	isMember, err := b.client.SIsMember(ctx, "participants:", participant.Username).Result()
	if err != nil {
		return false, err
	}
	if !isMember {
		return false, fmt.Errorf("participant not found %s", participant.Username)
	}

	password, err := b.client.HGet(ctx, participant.Username, "password").Result()
	if err != nil {
		return false, err
	}

	if passwordHash, _ := utils.SHA256(participant.Password); passwordHash != password {
		return false, nil
	}

	return true, nil
}

func (b *RedisBackend) HasChannel(ctx context.Context, channelName string) (bool, error) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	isMember, err := b.client.SIsMember(ctx, "channels:", channelName).Result()
	return isMember, err
}

func (b *RedisBackend) RegisterChannel(ctx context.Context, channel *apitypes.Channel) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	isMember, err := b.client.SIsMember(ctx, "channels:", channel.Name).Result()
	if err != nil {
		return err
	}

	if isMember {
		return fmt.Errorf("channel already exists")
	}

	if _, err := b.client.HSet(ctx, channel.Name, map[string]interface{}{
		"name":          channel.Name,
		"domain":        channel.Domain,
		"creator":       channel.Creator,
		"creation_time": channel.CreationTime,
	}).Result(); err != nil {
		return err
	}
	_, err = b.client.SAdd(ctx, "channels:", channel.Name).Result()
	return err
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
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	usernames, err := b.client.SMembers(ctx, "channels:").Result()
	if err != nil {
		return nil, err
	}

	result := &apitypes.ParticipanList{}
	if len(usernames) != 0 {
		participants := make([]*apitypes.Participant, len(usernames))
		_, err = b.client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			for _, username := range usernames {
				participant, err := pipe.HGetAll(ctx, username).Result()
				if err != nil {
					return err
				}

				joinTime, _ := time.Parse(time.Layout, participant["join_time"])
				participants = append(participants, &apitypes.Participant{
					Username:     participant["username"],
					Password:     participant["password"], // password is hashed
					EmailAddress: participant["email"],
					JoinTime:     joinTime,
				})
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
		result.Participants = participants
	}
	return result, nil
}

func (b *RedisBackend) StoreMessage(message *apitypes.ChatMessage) error {
	return nil
}
