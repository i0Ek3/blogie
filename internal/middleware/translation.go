package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		// also you can use different languages as you wish
		unicode := ut.New(en.New(), zh.New())
		locale := c.GetHeader("locale")
		// we cannot get translator while locale field non-exist in Header
		// so we need add if judge to checks if locale is empty
		if locale == "" {
			locale = "zh"
		}
		trans, _ := unicode.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
			default:
				_ = en_translations.RegisterDefaultTranslations(v, trans)
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}
