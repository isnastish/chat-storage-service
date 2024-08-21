package backend

import (
	"context"
	"github.com/isnastish/chat-backend/pkg/apitypes"
)

type StorageBackend interface {
	HasParticipant(context.Context, string) (bool, error)                      // experimental (RegisterParticipant should return an error if participant is already registered)
	RegisterParticipant(context.Context, *apitypes.Participant) error          // error is returned if participant doesn't exist
	AuthorizeParticipant(context.Context, *apitypes.Participant) (bool, error) // this should probably return an error instead if the authorization fails
	HasChannel(context.Context, string) (bool, error)                          // experimental (RegesterChannel should return an error if channel is already registered)
	RegisterChannel(context.Context, *apitypes.Channel) error                  // error is returned if channel doesn't exist
	DeleteChannel(context.Context, string) (bool, error)
	GetGeneralChatHistory(context.Context) (*apitypes.ChatHistory, error)
	GetChannelHistory(context.Context, string) (*apitypes.ChatHistory, error)
	GetChannelList(context.Context) (*apitypes.ChannelList, error)
	GetParticipantList(context.Context) (*apitypes.ParticipanList, error)
	StoreMessage(*apitypes.ChatMessage) error
}
