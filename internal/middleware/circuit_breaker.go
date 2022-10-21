package middleware

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

func CircuitBreaker() gin.HandlerFunc {
	hystrix.ConfigureCommand("blogie", hystrix.CommandConfig{
		Timeout:                10,
		MaxConcurrentRequests:  100,
		ErrorPercentThreshold:  25,
		RequestVolumeThreshold: 3,
		SleepWindow:            1000,
	})

	return func(c *gin.Context) {
		hystrix.Go("blogie", func() error {
			c.Next()

			return nil
		}, func(err error) error {
			c.Abort()

			return err
		})
	}
}
