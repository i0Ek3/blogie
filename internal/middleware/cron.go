package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/internal/model"
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron"
)

func Cron(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Starting...")
		cr := cron.New()
		cr.AddFunc("* * * * * *", func() {
			log.Println("Run model.CleanAllTag...")
			model.CleanAllTag(db)
		})
		cr.AddFunc("* * * * * *", func() {
			log.Println("Run model.CleanAllArticle...")
			model.CleanAllArticle(db)
		})
		cr.Start()
		t := time.NewTimer(time.Second * 10)
		for {
			select {
			case <-t.C:
				t.Reset(time.Second * 10)
			default:
				c.Next()
			}
		}
	}
}
