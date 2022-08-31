package cache

import (
	"github.com/i0Ek3/blogie/internal/model"
	"strconv"
	"strings"

	"github.com/i0Ek3/blogie/pkg/errcode"
)

type Article struct {
	model.Article
	TagID int

	PageNum  int
	PageSize int
}

func (a *Article) GetArticleKey() string {
	return errcode.CacheArticle + "_" + strconv.Itoa(int(a.ID))
}

func (a *Article) GetArticlesKey() string {
	keys := []string{
		errcode.CacheArticle,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(int(a.ID)))
	}
	if a.TagID > 0 {
		keys = append(keys, strconv.Itoa(a.TagID))
	}
	keys = append(keys, strconv.Itoa(int(a.State)))
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}
