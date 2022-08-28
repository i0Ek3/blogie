package middleware

import "github.com/gin-gonic/gin"

func CircuitBreaker() gin.HandlerFunc {
	return func(c *gin.Context){
        // TODO: https://github.com/go-kratos/aegis/blob/main/circuitbreaker/sre/sre.go
	    c.Next()
	}
}
