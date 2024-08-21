package backend

import "github.com/isnastish/chat-backend/pkg/apitypes"

type StorageBackend interface {
	HasParticipant(string) bool                      // experimental (RegisterParticipant should return an error if participant is already registered)
	RegisterParticipant(*apitypes.Participant) error // error is returned if participant doesn't exist
	AuthorizeParticipant(*apitypes.Participant) bool // this should probably return an error instead if the authorization fails
	HasChannel(string) bool                          // experimental (RegesterChannel should return an error if channel is already registered)
	RegisterChannel(*apitypes.Channel) error         // error is returned if channel doesn't exist
	DeleteChannel(string) bool
	GetGeneralChatHistory() *apitypes.ChatHistory
	GetChannelHistory(string) *apitypes.ChatHistory
	GetChannelList() *apitypes.ChannelList
	GetParticipantList() *apitypes.ParticipanList
	StoreMessage(*apitypes.ChatMessage)
}
