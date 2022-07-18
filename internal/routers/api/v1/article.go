package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/pkg/app"
	"github.com/i0Ek3/blogie/pkg/errcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary Get An Article
// @Produce json
// @Param id path int true "article id"
// @Success 200 {object} model.Article "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal server error"
// @Router /api/v1/articles/{id} [get]
func (t Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.InternalServerError)
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
func (t Article) List(c *gin.Context) {}

// @Summary Add A New Article
// @Produce json
// @Param tag_id body string true "tag id"
// @Param title body string true "article name"
// @Param desc body string false "description"
// @Param cover_image_url body string true "cover image url"
// @Param content body string true "article content"
// @Param created_by body string true "creator"
// @Param state body int false "state"
// @Success 200 {object} model.Article "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal server error"
// @Router /api/v1/articles [post]
func (t Article) Create(c *gin.Context) {}

// @Summary Update An Article
// @Produce json
// @Param tag_id body string false "tag id"
// @Param title body string false "article name"
// @Param desc body string false "description"
// @Param cover_image_url body string true "cover image url"
// @Param content body string true "article content"
// @Param modified_by body string true "updator"
// @Success 200 {object} model.Article "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal server error"
// @Router /api/v1/articles/{id} [put]
func (t Article) Update(c *gin.Context) {}

// @Summary Delete An Article
// @Produce json
// @Param id path int true "article id"
// @Success 200 {object} model.Article "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal server error"
// @Router /api/v1/articles/{id} [delete]
func (t Article) Delete(c *gin.Context) {}
