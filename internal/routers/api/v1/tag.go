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

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {
	debug.DebugHere("tag::", "Get")
	app.NewResponse(c).ToErrorResponse(errcode.InternalServerError)
}

// @Summary Get Tag List
// @Produce json
// @Param name query string false "tag name" maxlength(100)
// @Param state query int false "state" Enums(0,1) default(1)
// @Param page query int false "page"
// @Param page_size query int false "page size"
// @Success 200 {object} model.TagSwagger "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal server error"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	debug.DebugHere("tag::", "List")
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf(c, "svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
}

// @Summary Add A New Tag
// @Produce json
// @Param name body string true "tag name" minlength(3) maxlength(100)
// @Param state body int false "state" Enums(0,1) default(1)
// @Param created_by body string true "creator" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal server error"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	debug.DebugHere("tag::", "Create")
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreatTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary Update A Tag
// @Produce json
// @Param id path int true "tag id"
// @Param name body string false "tag name" minlength(3) maxlength(100)
// @Param state body int false "state" Enums(0,1) default(1)
// @Param modified_by body string true "updator" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal server error"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	debug.DebugHere("tag::", "Update")
	param := service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary Delete A Tag
// @Produce json
// @Param id path int true "tag id"
// @Success 200 {object} string "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal server error"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	debug.DebugHere("tag::", "Delete")
	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
}
