package service

import (
	"context"

	otg "github.com/eddycjy/opentracing-gorm"
	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otg.WithContext(svc.ctx, global.DBEngine))
	return svc
}
