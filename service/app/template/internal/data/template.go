package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"service/app/template/internal/biz"
)

type templateRepo struct {
	data *Data
	log  *log.Helper
}

// NewTemplateRepo .
func NewTemplateRepo(data *Data, logger log.Logger) biz.TemplateRepo {
	return &templateRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *templateRepo) CreateModel(ctx context.Context, g *biz.Model) error {
	return nil
}
