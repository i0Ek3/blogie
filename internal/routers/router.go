package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/internal/middleware"
	"github.com/i0Ek3/blogie/internal/routers/api"
	v1 "github.com/i0Ek3/blogie/internal/routers/api/v1"
	"github.com/i0Ek3/blogie/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/i0Ek3/blogie/docs"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.ContextTimeout * time.Second))
	r.Use(middleware.Tracing())
	r.Use(middleware.Translations())

	url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// upload file and access file on static address
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	// Notes: StaticFS() -> createStaticHandler() -> fileServer.ServerHTTP()
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.POST("/auth", api.GetAuth)

	// Group Router
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		tag := v1.NewTag()
		{
			apiv1.POST("/tags", tag.Create)
			apiv1.DELETE("/tags/:id", tag.Delete)
			apiv1.PUT("/tags/:id", tag.Update)
			apiv1.PATCH("/tags/:id/state", tag.Update)
			apiv1.GET("/tags", tag.List)
		}

		article := v1.NewArticle()
		{
			apiv1.POST("/articles", article.Create)
			apiv1.DELETE("/articles/:id", article.Delete)
			apiv1.PUT("/articles/:id", article.Update)
			apiv1.PATCH("/articles/:id/state", article.Update)
			apiv1.GET("/articles/:id", article.Get)
			apiv1.GET("/articles", article.List)
		}
	}

	return r
}
