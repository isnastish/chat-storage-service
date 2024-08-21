package redis

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/isnastish/chat-backend/pkg/apitypes"
	"github.com/isnastish/chat-backend/pkg/testsetup"
)

const redisPort = 6379

func TestMain(m *testing.M) {
	var tearDown bool
	var exitCode = 0

	defer func() {
		if tearDown {
			testsetup.KillRedisEmulator()
		}
		os.Exit(exitCode)
	}()

	tearDown, err := testsetup.StartRedisEmulator(redisPort)
	if err == nil {
		exitCode = m.Run()
	}
}

func TestRegisterParticipant(t *testing.T) {
	rb, err := NewRedisBackend(&apitypes.RedisConfig{
		Endpoint: fmt.Sprintf("localhost:%d", redisPort),
		Username: "",
		Password: "12345",
	})
	if err != nil {
		t.Fatalf("Failed to init backend %v", err)
	}

	rb.RegisterParticipant(context.Background(), &apitypes.Participant{})
}
