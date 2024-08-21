package service

import (
	"github.com/isnastish/chat-backend/pkg/api/backend"
	"github.com/isnastish/chat-backend/pkg/apitypes"
)

type Service struct {
	backend           backend.StorageBackend
	AllowUnauthorized bool
}

// experimental (RegisterParticipant should return an error if participant is already registered)
func (s *Service) HasParticipant(string) bool {
	return false
}

func (s *Service) RegisterParticipant(*apitypes.Participant) {

}

func (s *Service) AuthorizeParticipant(*apitypes.Participant) bool {
	return false
}

// experimental (RegesterChannel should return an error if channel is already registered)
func (s *Service) HasChannel(string) bool {
	return false
}

func (s *Service) RegisterChannel(*apitypes.Channel) {

}

func (s *Service) DeleteChannel(string) bool {
	// return true if a channel was deleted
	return false
}

func (s *Service) GetGeneralChatHistory() *apitypes.ChatHistory {
	return &apitypes.ChatHistory{}
}

func (s *Service) GetChannelHistory(string) {

}

func (s *Service) GetChannelList() *apitypes.ChannelList {
	return &apitypes.ChannelList{}
}

func (s *Service) GetParticipantList() *apitypes.ParticipanList {
	return &apitypes.ParticipanList{}
}

func (s *Service) StoreMessage(*apitypes.ChatMessage) {
}
