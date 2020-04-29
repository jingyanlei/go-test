package dao

import (
	"fmt"

	"github.com/go-redis/redis"
)

// var ProviderRedis = wire.NewSet(NewRedis)

func NewRedis() (r *redis.Client, cf func(), err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println("redis ping", pong, err)
	cf = func() { fmt.Println("redis close"); client.Close() }
	// _ = cf
	return
}
