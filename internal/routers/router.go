package routers

import (
	"net/http"
	"time"

	v1 "github.com/i0Ek3/blogie/internal/routers/api/v1"
	"github.com/i0Ek3/blogie/pkg/debug"
	"github.com/i0Ek3/blogie/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	_ "github.com/i0Ek3/blogie/docs"
	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/internal/middleware"
	"github.com/i0Ek3/blogie/internal/routers/api"
)

var (
	limiters = limiter.NewMethodLimiter().AddBuckets(limiter.BucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	})
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" && !global.EnableSetting.Enable {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

    d := global.AppSetting.ContextTimeout * time.Second
	r.Use(middleware.ContextTimeout(d))
	r.Use(middleware.RateLimiter(limiters))
	r.Use(middleware.CircuitBreaker())
	r.Use(middleware.Cors())
	r.Use(middleware.AppInfo())
	r.Use(middleware.Tracing())
	r.Use(middleware.Translations())

	r.GET("debug/vars", api.Expvar)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Upload file and access file on static address
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)

	// Setting up file services to provide access to static resources
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.POST("/auth", api.GetAuth)

	tag := v1.NewTag()
	article := v1.NewArticle()

	apiv1 := r.Group("api/v1")
	apiv1.Use(middleware.JWT(), middleware.Cron(global.GDB))
	{
		tags := apiv1.Group("/tags")
		{
			debug.DebugHere("tag", "POST::Create")
			tags.POST("", tag.Create)
			debug.DebugHere("tag", "DELETE::Delete")
			tags.DELETE(":id", tag.Delete)
			debug.DebugHere("tag", "PUT::Update")
			tags.PUT(":id", tag.Update)
			debug.DebugHere("tag", "PATCH::Update")
			tags.PATCH(":id/state", tag.Update)
			debug.DebugHere("tag", "GET::Get")
			tags.GET(":id", tag.Get)
			debug.DebugHere("tag", "GET::List")
			tags.GET("", tag.List)
		}

		articles := apiv1.Group("/articles")
		{
			debug.DebugHere("article", "POST::Create")
			articles.POST("", article.Create)
			debug.DebugHere("article", "DELETE::Delete")
			articles.DELETE(":id", article.Delete)
			debug.DebugHere("article", "PUT::Update")
			articles.PUT(":id", article.Update)
			debug.DebugHere("article", "PATCH::Update")
			articles.PATCH(":id/state", article.Update)
			debug.DebugHere("article", "GET::Get")
			articles.GET(":id", article.Get)
			debug.DebugHere("article", "GET::List")
			articles.GET("", article.List)
		}
	}

	return r
}
