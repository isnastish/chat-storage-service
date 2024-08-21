package testsetup

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"

	"github.com/isnastish/chat-backend/pkg/log"
)

// Start redis instance inside a docker container on port `port`.
func StartRedisEmulator(port uint) (bool, error) {
	tearDown := false

	cmd := exec.Command("docker", "run", "--rm", "--name", "redis-emulator", "-p", fmt.Sprintf("%d:6379", port), "redis:latest")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return tearDown, err
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		return tearDown, err
	}

	tearDown = true

	var strBuilder strings.Builder
	timer := time.NewTimer(3 * time.Minute)
	buf := make([]byte, 256)

	for {
		select {
		case <-timer.C:
			return tearDown, err
		default:
		}
		n, err := stdout.Read(buf[:])
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Logger.Fatal("Failed to read stdout %v", err)
		}
		if n > 0 {
			stdoutStr := string(buf[:n])
			log.Logger.Info(stdoutStr)
			strBuilder.WriteString(stdoutStr)
			if strings.Contains(strBuilder.String(), "Ready to accept connections") {
				break
			}
		}
	}
	return tearDown, err
}

// Kill docker container running redis instance
func KillRedisEmulator() {
	cmd := exec.Command("docker", "rm", "--force", "redis-emulator")
	err := cmd.Run()
	if err != nil {
		log.Logger.Error("Failed to tear down redis container %s", err)
	} else {
		log.Logger.Info("Redis container shut down")
	}
}
