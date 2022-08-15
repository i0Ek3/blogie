package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/internal/service"
	"github.com/i0Ek3/blogie/pkg/app"
	"github.com/i0Ek3/blogie/pkg/convert"
	"github.com/i0Ek3/blogie/pkg/debug"
	"github.com/i0Ek3/blogie/pkg/errcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary Get An Article
// @Produce json
// @Param id path int true "article id"
// @Success 200 {object} model.ArticleSwagger "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal server error"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	debug.DebugHere("article::", "Get")
	param := service.ArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	article, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(article)
}

// @Summary Get Article List
// @Produce json
// @Param name query string false "article name"
// @Param tag_id query int false "tag id"
// @Param state query int false "state"
// @Param page query int false "page"
// @Param page_size query int false "page size"
// @Success 200 {object} model.ArticleSwagger "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal server error"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	debug.DebugHere("article::", "List")
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	articles, totalRows, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}

	response.ToResponseList(articles, totalRows)
}

// @Summary Add A New Article
// @Produce json
// @Param tag_id body string true "tag id"
// @Param title body string true "article name"
// @Param desc body string false "description"
// @Param cover_image_url body string true "cover image url"
// @Param content body string true "article content"
// @Param created_by body string true "creator"
// @Param state body int false "state"
// @Success 200 {object} model.ArticleSwagger "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal server error"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	debug.DebugHere("article::", "Create")
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary Update An Article
// @Produce json
// @Param tag_id body string false "tag id"
// @Param title body string false "article name"
// @Param desc body string false "description"
// @Param cover_image_url body string true "cover image url"
// @Param content body string true "article content"
// @Param modified_by body string true "updator"
// @Success 200 {object} model.ArticleSwagger "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal server error"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {
	debug.DebugHere("article::", "Update")
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary Delete An Article
// @Produce json
// @Param id path int true "article id"
// @Success 200 {object} string "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal server error"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	debug.DebugHere("article::", "Delete")
	param := service.DeleteArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}
