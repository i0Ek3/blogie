package middleware

import (
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func Redis() gin.HandlerFunc {
	redisStore := persist.NewRedisStore(redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	}))

	return func(c *gin.Context) {
		cache.CacheByRequestURI(redisStore, 2*time.Second)
		c.Next()
	}
}
