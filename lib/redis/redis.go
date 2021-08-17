package redis

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func Client() *redis.Client {
	opt, _ := redis.ParseURL(os.Getenv("REDIS"))

	return redis.NewClient(opt)
}
