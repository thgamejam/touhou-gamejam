package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"service/app/account/internal/biz"
)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

// NewAccountRepo .
func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *accountRepo) CreateModel(ctx context.Context, g *biz.Model) error {
	return nil
}
