package database

import(
	"e/config"
	"fmt"
	"github.com/go-redis/redis/v8"

)

func ConnectionRedisDb(config *config.Config) *redis.Client{
	rdb:= redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})

	fmt.Println("Connected Successfully to the database (Redis)")
	return rdb
}