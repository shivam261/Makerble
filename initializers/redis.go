package initializers

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func ConnectToRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),     // e.g. "localhost:6379"
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

	// Test the connection
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Println("Could not connect to Redis: %v", err)
	} else {
		log.Println("Connected to Redis successfully")
	}

	// Set a key with an expiration time
	err = RedisClient.Set(Ctx, "test_key", "test_value", 10*time.Minute).Err()
	if err != nil {
		log.Println("Could not set test key in Redis: %v", err)
	}
}
