package utils

import (
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

func RedisCache() *cache.Cache {
	redisDB := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
		},
	})

	pillshareCache := cache.New(&cache.Options{
		Redis:      redisDB,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return pillshareCache

}
