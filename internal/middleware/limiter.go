package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/pkg/app"
	"github.com/i0Ek3/blogie/pkg/errcode"
	"github.com/i0Ek3/blogie/pkg/limiter"
)

func RateLimiter(l limiter.BaseLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			// TakeAvailable returns the available buckets
			count := bucket.TakeAvailable(1)
			// if available buckets' count equals 0, which means there are too many requests
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
