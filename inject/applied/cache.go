package applied

import (
	"demo/internal/applied/cache"
	"demo/internal/applied/cache/redis"
)

func InitCacheConnection() cache.Cache {
	return redis.NewConnection()
}
