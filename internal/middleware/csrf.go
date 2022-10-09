package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	wrap "github.com/turtlemonvh/gin-wraphh"
)

func CSRF(authKey []byte, domain string) gin.HandlerFunc {
	return func(c *gin.Context) {
		wrap.WrapHH(csrf.Protect(authKey,
			csrf.Secure(false),
			csrf.HttpOnly(true),
			csrf.Domain(domain),
			csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusForbidden)
				_, _ = w.Write([]byte(`{"message": "CSRF token invalid"}`))
			})),
		))
		c.Next()
	}
}
