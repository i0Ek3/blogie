package app

import (
	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/pkg/convert"
)

// GetPage gets the page according page field in the DB
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

// GetPageSize gets the page size according page_size field in the DB
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

// GetPageOffset gets the offset according page and pageSize
func GetPageOffset(page, pageSize int) int {
	offset := 0
	if page > 0 {
		offset = (page - 1) * pageSize
	}

	return offset
}
