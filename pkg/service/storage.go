package service

import (
	"context"
	"fmt"

	"github.com/isnastish/chat-backend/pkg/apitypes"
	"github.com/isnastish/chat-backend/pkg/backend"
	"github.com/isnastish/chat-backend/pkg/backend/redis"
)

type Service struct {
	backend           backend.StorageBackend
	AllowUnauthorized bool
}

func NewService(config *apitypes.Config) (*Service, error) {
	var backend backend.StorageBackend
	var err error

	switch config.Backend {
	case apitypes.RedisBackend:
		backend, err = redis.NewRedisBackend(config.RedisConfig)
		if err != nil {
			return nil, fmt.Errorf("Failed to init redis backend %v", err)
		}
	}

	return &Service{
		AllowUnauthorized: config.AllowUnauthorized,
		backend:           backend,
	}, nil
}

// experimental (RegisterParticipant should return an error if participant is already registered)
func (s *Service) HasParticipant(ctx context.Context, participantName string) (bool, error) {
	return s.backend.HasParticipant(ctx, participantName)
}

func (s *Service) RegisterParticipant(ctx context.Context, participant *apitypes.Participant) error {
	return s.backend.RegisterParticipant(ctx, participant)
}

func (s *Service) AuthorizeParticipant(ctx context.Context, participant *apitypes.Participant) (bool, error) {
	return s.backend.AuthorizeParticipant(ctx, participant)
}

// experimental (RegesterChannel should return an error if channel is already registered)
func (s *Service) HasChannel(ctx context.Context, channelName string) (bool, error) {
	return s.backend.HasChannel(ctx, channelName)
}

func (s *Service) RegisterChannel(ctx context.Context, channel *apitypes.Channel) error {
	return s.backend.RegisterChannel(ctx, channel)
}

func (s *Service) DeleteChannel(ctx context.Context, channelName string) (bool, error) {
	// return true if a channel was deleted
	return s.backend.DeleteChannel(ctx, channelName)
}

func (s *Service) GetGeneralChatHistory(ctx context.Context) (*apitypes.ChatHistory, error) {
	return s.backend.GetGeneralChatHistory(ctx)
}

func (s *Service) GetChannelHistory(ctx context.Context, channelName string) (*apitypes.ChatHistory, error) {
	return s.backend.GetChannelHistory(ctx, channelName)
}

func (s *Service) GetChannelList(ctx context.Context) (*apitypes.ChannelList, error) {
	return s.backend.GetChannelList(ctx)
}

func (s *Service) GetParticipantList(ctx context.Context) (*apitypes.ParticipanList, error) {
	return s.backend.GetParticipantList(ctx)
}

func (s *Service) StoreMessage(messages *apitypes.ChatMessage) {
	s.backend.StoreMessage(messages)
}
