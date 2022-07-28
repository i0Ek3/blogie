package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/internal/middleware"
	"github.com/i0Ek3/blogie/internal/routers/api"
	v1 "github.com/i0Ek3/blogie/internal/routers/api/v1"
	"github.com/i0Ek3/blogie/pkg/debug"
	"github.com/i0Ek3/blogie/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/i0Ek3/blogie/docs"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.BucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" && !global.EnableSetting.Enable {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.ContextTimeout * time.Second))
	r.Use(middleware.AppInfo())
	r.Use(middleware.Tracing())
	r.Use(middleware.Translations())

	r.GET("debug/vars", api.Expvar)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// upload file and access file on static address
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	// NOTES: StaticFS() -> createStaticHandler() -> fileServer.ServerHTTP()
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.POST("/auth", api.GetAuth)

	// group router
	apiv1 := r.Group("api/v1")
	apiv1.Use(middleware.JWT())
	{
		tag := v1.NewTag()
		tags := apiv1.Group("/tags")
		{
			debug.DebugHere("tag", "tag.Create")
			tags.POST("", tag.Create)
			debug.DebugHere("tag", "tag.Delete")
			tags.DELETE(":id", tag.Delete)
			debug.DebugHere("tag", "tag.Update")
			tags.PUT(":id", tag.Update)
			debug.DebugHere("tag", "tag.Update")
			tags.PATCH(":id/state", tag.Update)
			debug.DebugHere("tag", "tag.List")
			tags.GET("", tag.List)
		}

		article := v1.NewArticle()
		articles := apiv1.Group("/articles")
		{
			debug.DebugHere("article", article)
			articles.POST("", article.Create)
			articles.DELETE(":id", article.Delete)
			articles.PUT(":id", article.Update)
			articles.PATCH(":id/state", article.Update)
			articles.GET(":id", article.Get)
			articles.GET("", article.List)
		}
	}

	return r
}
