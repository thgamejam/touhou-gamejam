package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUseCase)

type UserRepo interface {
	// GetUserByAccountID 通过账户ID获取用户
	GetUserByAccountID(ctx context.Context, accountID uint32) (user *UserInfo, err error)
	//CreateUser 创建用户
	CreateUser(ctx context.Context, accountID uint32) (user *UserInfo, err error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}
