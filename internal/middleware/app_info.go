package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/pkg/version"
)

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", version.AppName)
		c.Set("app_version", version.Version)
		c.Next()
	}
}
