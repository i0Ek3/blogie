package service

import (
	"context"

	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
