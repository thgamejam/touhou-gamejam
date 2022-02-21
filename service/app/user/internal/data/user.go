package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"service/app/user/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateModel(ctx context.Context, g *biz.Model) error {
	return nil
}
