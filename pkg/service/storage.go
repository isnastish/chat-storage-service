package service

import (
	"github.com/isnastish/chat-backend/pkg/apitypes"
	"github.com/isnastish/chat-backend/pkg/backend"
)

type Service struct {
	backend           backend.StorageBackend
	AllowUnauthorized bool
}

// experimental (RegisterParticipant should return an error if participant is already registered)
func (s *Service) HasParticipant(participantName string) bool {
	return s.backend.HasParticipant(participantName)
}

func (s *Service) RegisterParticipant(participant *apitypes.Participant) error {
	return s.backend.RegisterParticipant(participant)
}

func (s *Service) AuthorizeParticipant(participant *apitypes.Participant) bool {
	return s.backend.AuthorizeParticipant(participant)
}

// experimental (RegesterChannel should return an error if channel is already registered)
func (s *Service) HasChannel(channelName string) bool {
	return s.backend.HasChannel(channelName)
}

func (s *Service) RegisterChannel(channel *apitypes.Channel) error {
	return s.backend.RegisterChannel(channel)
}

func (s *Service) DeleteChannel(channelName string) bool {
	// return true if a channel was deleted
	return s.backend.DeleteChannel(channelName)
}

func (s *Service) GetGeneralChatHistory() *apitypes.ChatHistory {
	return &apitypes.ChatHistory{}
}

func (s *Service) GetChannelHistory(channelName string) *apitypes.ChatHistory {
	return s.backend.GetChannelHistory(channelName)
}

func (s *Service) GetChannelList() *apitypes.ChannelList {
	return s.backend.GetChannelList()
}

func (s *Service) GetParticipantList() *apitypes.ParticipanList {
	return s.backend.GetParticipantList()
}

func (s *Service) StoreMessage(messages *apitypes.ChatMessage) {
	s.backend.StoreMessage(messages)
}
