package service

import (
	"context"

	"github.com/isnastish/chat-backend/pkg/apitypes"
	"github.com/isnastish/chat-backend/pkg/backend"
)

type Service struct {
	backend           backend.StorageBackend
	AllowUnauthorized bool
}

// experimental (RegisterParticipant should return an error if participant is already registered)
func (s *Service) HasParticipant(ctx context.Context, participantName string) bool {
	return s.backend.HasParticipant(ctx, participantName)
}

func (s *Service) RegisterParticipant(ctx context.Context, participant *apitypes.Participant) error {
	return s.backend.RegisterParticipant(ctx, participant)
}

func (s *Service) AuthorizeParticipant(ctx context.Context, participant *apitypes.Participant) bool {
	return s.backend.AuthorizeParticipant(ctx, participant)
}

// experimental (RegesterChannel should return an error if channel is already registered)
func (s *Service) HasChannel(ctx context.Context, channelName string) bool {
	return s.backend.HasChannel(ctx, channelName)
}

func (s *Service) RegisterChannel(ctx context.Context, channel *apitypes.Channel) error {
	return s.backend.RegisterChannel(ctx, channel)
}

func (s *Service) DeleteChannel(ctx context.Context, channelName string) bool {
	// return true if a channel was deleted
	return s.backend.DeleteChannel(ctx, channelName)
}

func (s *Service) GetGeneralChatHistory(ctx context.Context) *apitypes.ChatHistory {
	return s.backend.GetGeneralChatHistory(ctx)
}

func (s *Service) GetChannelHistory(ctx context.Context, channelName string) *apitypes.ChatHistory {
	return s.backend.GetChannelHistory(ctx, channelName)
}

func (s *Service) GetChannelList(ctx context.Context) *apitypes.ChannelList {
	return s.backend.GetChannelList(ctx)
}

func (s *Service) GetParticipantList(ctx context.Context) *apitypes.ParticipanList {
	return s.backend.GetParticipantList(ctx)
}

func (s *Service) StoreMessage(messages *apitypes.ChatMessage) {
	s.backend.StoreMessage(messages)
}
