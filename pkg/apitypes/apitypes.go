package apitypes

import (
	"time"
)

type Backend int

const (
	_ Backend = iota
	MemoryBackend
	RedisBackend
	FaunaBackend
	DynamoBackend
)

type Config struct {
	Backend            Backend
	RedisConfig        *RedisConfig
	DynamoConfig       *DynamoConfig
	FaunaConfig        *FaunaConfig
	ParticipantTimeout time.Duration
	AllowUnauthorized  bool
}

type RedisConfig struct {
	Username string
	Password string
	Endpoint string
}

type FaunaConfig struct {
	FaunaSecret    string
	FaunaEndpoinrt string
}

type DynamoConfig struct {
	AccessKeyID string
	SecretKey   string
}

type ChatMessage struct {
	Contents string
	Sender   string
	Channel  *string // optional
	SendTime time.Time
}

type Channel struct {
	Name         string
	Domain       string
	Creator      string
	CreationTime time.Time
}

type Participant struct {
	Username     string
	Password     string
	EmailAddress string
	JoinTime     time.Time
}

type ChatHistory struct {
	Messages []*ChatMessage
}

type ChannelList struct {
	Channels []*Channel
}

type ParticipanList struct {
	Participants []*Participant
}
