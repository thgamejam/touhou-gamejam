package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Model struct {
	Hello string
}

type PassportRepo interface {
	CreateModel(context.Context, *Model) error
}

type PassportUseCase struct {
	repo PassportRepo
	log  *log.Helper
}

func NewPassportUseCase(repo PassportRepo, logger log.Logger) *PassportUseCase {
	return &PassportUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PassportUseCase) Create(ctx context.Context, g *Model) error {
	return uc.repo.CreateModel(ctx, g)
}
