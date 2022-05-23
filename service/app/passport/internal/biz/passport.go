package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Model struct {
	Hello string
}

type Account struct {
	Email    string
	Password string
	Hash     string
}

type PassportRepo interface {
	// PrepareCreateAccount 预创建账户到缓存
	PrepareCreateAccount(ctx context.Context, account Account) error
}

type PassportUseCase struct {
	repo PassportRepo
	log  *log.Helper
}

func NewPassportUseCase(repo PassportRepo, logger log.Logger) *PassportUseCase {
	return &PassportUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PassportUseCase) PrepareCreateAccount(ctx context.Context, account Account, token string) error {

	// TODO 检测验证码token

	return uc.repo.PrepareCreateAccount(ctx, account)
}
