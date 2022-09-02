package middleware

import "github.com/gin-gonic/gin"

func CircuitBreaker() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
		c.Next()
	}
}
