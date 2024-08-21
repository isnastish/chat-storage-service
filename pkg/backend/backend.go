package backend

import (
	"context"
	"github.com/isnastish/chat-backend/pkg/apitypes"
)

type StorageBackend interface {
	HasParticipant(context.Context, string) bool                      // experimental (RegisterParticipant should return an error if participant is already registered)
	RegisterParticipant(context.Context, *apitypes.Participant) error // error is returned if participant doesn't exist
	AuthorizeParticipant(context.Context, *apitypes.Participant) bool // this should probably return an error instead if the authorization fails
	HasChannel(context.Context, string) bool                          // experimental (RegesterChannel should return an error if channel is already registered)
	RegisterChannel(context.Context, *apitypes.Channel) error         // error is returned if channel doesn't exist
	DeleteChannel(context.Context, string) bool
	GetGeneralChatHistory(context.Context) *apitypes.ChatHistory
	GetChannelHistory(context.Context, string) *apitypes.ChatHistory
	GetChannelList(context.Context) *apitypes.ChannelList
	GetParticipantList(context.Context) *apitypes.ParticipanList
	StoreMessage(*apitypes.ChatMessage)
}
