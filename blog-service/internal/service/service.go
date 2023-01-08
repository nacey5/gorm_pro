package service

import (
	"golang.org/x/net/context"
	"gorm_pro/blog-service/global"
	"gorm_pro/blog-service/internal/dao"
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
