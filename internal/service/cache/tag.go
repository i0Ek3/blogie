package cache

import (
	"github.com/i0Ek3/blogie/internal/model"
	"strconv"
	"strings"

	"github.com/i0Ek3/blogie/pkg/errcode"
)

type Tag struct {
	model.Tag

	PageNum  int
	PageSize int
}

func (t *Tag) GetTagsKey() string {
	keys := []string{
		errcode.CacheTag,
		"LIST",
	}

	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	keys = append(keys, strconv.Itoa(int(t.State)))

	if t.PageNum > 0 {
		keys = append(keys, strconv.Itoa(t.PageNum))
	}
	if t.PageSize > 0 {
		keys = append(keys, strconv.Itoa(t.PageSize))
	}

	return strings.Join(keys, "_")
}
