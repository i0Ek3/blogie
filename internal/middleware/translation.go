package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/global"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.GetHeader("locale")
		trans, found := global.Ut.GetTranslator(locale)
		if found {
			c.Set("trans", trans)
		} else {
			enTrans, _ := global.Ut.GetTranslator("en")
			c.Set("trans", enTrans)
		}
		c.Next()
	}
}
