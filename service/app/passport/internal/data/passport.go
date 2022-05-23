package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"service/app/passport/internal/biz"
)

type passportRepo struct {
	data *Data
	log  *log.Helper
}

// NewPassportRepo .
func NewPassportRepo(data *Data, logger log.Logger) biz.PassportRepo {
	return &passportRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *passportRepo) CreateModel(ctx context.Context, g *biz.Model) error {
	return nil
}
