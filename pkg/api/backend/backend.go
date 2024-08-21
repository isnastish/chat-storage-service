package backend

import "github.com/isnastish/chat-backend/pkg/apitypes"

type StorageBackend interface {
	HasParticipant(string) bool // experimental (RegisterParticipant should return an error if participant is already registered)
	RegisterParticipant(*apitypes.Participant)
	AuthorizeParticipant(*apitypes.Participant) bool
	HasChannel(string) bool // experimental (RegesterChannel should return an error if channel is already registered)
	RegisterChannel(*apitypes.Channel)
	DeleteChannel(string) bool
	GetGeneralChatHistory() *apitypes.ChatHistory
	GetChannelHistory(string)
	GetChannelList() *apitypes.ChannelList
	GetParticipantList() *apitypes.ParticipanList
	StoreMessage(*apitypes.ChatMessage)
}
