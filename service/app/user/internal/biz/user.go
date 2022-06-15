package biz

import (
	"context"
)

type UserInfo struct {
	Name      string
	AvatarUrl string
	WorkCount uint32
	FansCount uint32
	Tags      []string
}

// GetUserByAccountID 通过账户ID获取用户信息
func (uc *UserUseCase) GetUserByAccountID(ctx context.Context, accountID uint32) (*UserInfo, error) {
	return uc.repo.GetUserByAccountID(ctx, accountID)
}

// CreateUser 根据账户ID创建用户
func (uc *UserUseCase) CreateUser(ctx context.Context, accountID uint32) (userInfo *UserInfo, err error) {
	return uc.repo.CreateUser(ctx, accountID)
}
