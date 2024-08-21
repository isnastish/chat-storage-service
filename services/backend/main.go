package main

import (
	"flag"
	"os"
	"strings"

	"github.com/isnastish/chat-backend/pkg/apitypes"
	"github.com/isnastish/chat-backend/pkg/log"
)

func main() {
	var config apitypes.Config

	backend := flag.String("backend", "", "Storage backend")
	grpcPort := flag.Uint("grpc_port", 50051, "Grpc server listening port")
	logLevel := flag.String("log_level", "info", "Log level")
	redisUser := flag.String("redis_user", "", "Redis username")
	redisEndpoint := flag.String("redis_endpoint", "localhost:6379", "Redis endpoint")
	redisPassword := flag.String("redis_password", "", "Redis password")

	flag.Parse()

	log.SetupGlobalLogLevel(*logLevel)

	*backend = strings.ToLower(*backend)
	switch *backend {
	case "memory":
		config.Backend = apitypes.MemoryBackend

	case "fauna":
		config.Backend = apitypes.FaunaBackend

	case "dynamo":
		config.Backend = apitypes.DynamoBackend

	case "redis":
		config.Backend = apitypes.RedisBackend
		config.RedisConfig = &apitypes.RedisConfig{
			Username: *redisUser,
			Password: *redisPassword,
			Endpoint: *redisEndpoint,
		}

	default:
		log.Logger.Fatal("Unknown backend %s", *backend)
	}

	_ = grpcPort

	os.Exit(0)
}
