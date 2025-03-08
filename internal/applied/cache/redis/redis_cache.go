package redis

import (
	"github.com/go-redis/redis"

	"mini/config"
	"mini/internal/applied/cache"
)

var (
	CacheConnection = new(redis.Client)
)

type Cache struct {
}

func NewConnection() cache.Cache {
	return &Cache{}
}

func (c Cache) ConnectCache() {
	CacheConnection = redis.NewClient(&redis.Options{
		Addr:     config.GetRedisAddress(),
		Password: config.C.Redis.Password,
		DB:       config.C.Redis.Db,
	})

	_, err := CacheConnection.Ping().Result()
	if err != nil {
		panic("failed to connect cache")
	}
}
