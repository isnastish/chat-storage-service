package redis

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/isnastish/chat-backend/pkg/apitypes"
	"github.com/isnastish/chat-backend/pkg/testsetup"
)

const RUN_REDIS_EMULATOR = false

const redisPort = 6379

func TestMain(m *testing.M) {
	var tearDown bool
	var err error
	var exitCode = 0

	defer func() {
		if RUN_REDIS_EMULATOR && tearDown {
			testsetup.KillRedisEmulator()
		}
		os.Exit(exitCode)
	}()

	if RUN_REDIS_EMULATOR {
		tearDown, err = testsetup.StartRedisEmulator(redisPort)
	}

	if err == nil {
		exitCode = m.Run()
	}
}

func GetRedisBackend(t *testing.T) *RedisBackend {
	rb, err := NewRedisBackend(&apitypes.RedisConfig{
		Endpoint: fmt.Sprintf("localhost:%d", redisPort),
		Username: "",
		Password: "12345",
	})
	if err != nil {
		t.Fatalf("Failed to init backend %v", err)
	}
	return rb
}

func TestRegisterParticipant(t *testing.T) {
	rb := GetRedisBackend(t)

	err := rb.RegisterParticipant(context.Background(), &apitypes.Participant{
		Username:     "saml",
		Password:     "1234@saml",
		EmailAddress: "saml@gmail.com",
		JoinTime:     time.Now().Add(2 * time.Hour),
	})
	if err != nil {
		t.Fatalf("Failed to register participant %v", err)
	}
}

func TestRegisterChannel(t *testing.T) {
	rb := GetRedisBackend(t)

	err := rb.RegisterChannel(context.Background(), &apitypes.Channel{
		Name:         "books",
		Domain:       "Channel for selling books",
		Creator:      "saml",
		CreationTime: time.Now(),
	})
	if err != nil {
		t.Fatalf("Failed to register channel %v", err)
	}
}

func TestAuthorizeParticipant(t *testing.T) {
	rb := GetRedisBackend(t)
	participant := &apitypes.Participant{
		Username:     "fedor",
		Password:     "fedor1234@__",
		EmailAddress: "fedor@gmil.com",
		JoinTime:     time.Now(),
	}

	_ = rb.RegisterParticipant(context.Background(), participant)

	authorized, err := rb.AuthorizeParticipant(context.Background(), participant)
	if err != nil {
		t.Fatalf("Call failed %v", err)
	}
	if !authorized {
		t.Fatalf("Failed to authorized participant")
	}
}
