package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Model struct {
	Hello string
}

type TemplateRepo interface {
	CreateModel(context.Context, *Model) error
}

type TemplateUseCase struct {
	repo TemplateRepo
	log  *log.Helper
}

func NewTemplateUseCase(repo TemplateRepo, logger log.Logger) *TemplateUseCase {
	return &TemplateUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TemplateUseCase) Create(ctx context.Context, g *Model) error {
	return uc.repo.CreateModel(ctx, g)
}
