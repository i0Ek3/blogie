package middleware

import (
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

func CircuitBreaker() gin.HandlerFunc {
	return func(c *gin.Context) {
		hystrix.ConfigureCommand("circuit breaker", hystrix.CommandConfig{
			Timeout:               int(10 * time.Second),
			MaxConcurrentRequests: 100,
			ErrorPercentThreshold: 25,
		})

		hystrix.Go("circuit breaker", func() error {
			c.Next()
			return nil
		}, func(err error) error {
			return err
		})
	}
}
