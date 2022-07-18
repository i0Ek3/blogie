package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/internal/middleware"
	"github.com/i0Ek3/blogie/internal/routers/api"
	v1 "github.com/i0Ek3/blogie/internal/routers/api/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/i0Ek3/blogie/docs"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())

	url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	article := v1.NewArticle()
	tag := v1.NewTag()

	// upload file and access file on static address
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	// Notes: StaticFS() -> createStaticHandler() -> fileServer.ServerHTTP()
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	// Group Router
	apiv1 := r.Group("/api/v1")
	{
		// API for tag
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		// API for article
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}
