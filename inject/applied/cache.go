package applied

import (
	"mini/internal/applied/cache"
	"mini/internal/applied/cache/redis"
)

func InitCacheConnection() cache.Cache {
	return redis.NewConnection()
}
