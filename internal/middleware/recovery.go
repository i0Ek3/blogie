package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/pkg/app"
	"github.com/i0Ek3/blogie/pkg/email"
	"github.com/i0Ek3/blogie/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != any(nil) {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)

				if global.EnableSetting.Enable {
					err := defaultMailer.SendMailTo(
						global.EmailSetting.To,
						fmt.Sprintf("panic appeared at: %d", time.Now().Unix()),
						fmt.Sprintf("error message: %v", err),
					)
					if err != nil {
						global.Logger.Panicf(c, "mail.SendMail err: %v", err)
					}
				}

				app.NewResponse(c).ToErrorResponse(errcode.InternalServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
