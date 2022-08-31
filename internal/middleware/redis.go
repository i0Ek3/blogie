package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/i0Ek3/blogie/pkg/redis"
)

func Redis() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := redis.Setup()
		if err != nil {
			return 
		}
		c.Next()
	}
}
