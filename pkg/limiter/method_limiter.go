package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// MethodLimiter defines a rate limiter to limit part of the routing
type MethodLimiter struct {
	*Limiter
	BaseLimiter
}

func NewMethodLimiter() BaseLimiter {
	return MethodLimiter{
		Limiter: &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)},
	}
}

// Key returns the key-value name according URI
func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}

	return uri[:index]
}

func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]

	return bucket, ok
}

func (l MethodLimiter) AddBuckets(rules ...BucketRule) BaseLimiter {
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			l.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}

	return l
}
