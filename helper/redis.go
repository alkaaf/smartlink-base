package helper

import (
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"os"
)

var c *redis.Client

func PGetRedis() *redis.Client {
	if c == nil {
		c = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASS"),
			DB:       0,
		})
	}
	return c
}

func GetRedis() *redis.Client {
	godotenv.Load()
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
}
