package service

import (
	"context"

	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/internal/dao"
	otg "github.com/smacker/opentracing-gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otg.SetSpanToGorm(svc.ctx, global.GDB))

	return svc
}
