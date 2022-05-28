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
	// CreatAccount 创建用户
	CreatAccount(ctx context.Context, sid string, key string) (id uint32, err error)

	// SignLoginToken 签署登录token
	SignLoginToken(ctx context.Context, accountID uint32) (token string, err error)

	// GetPublicKey 获取公钥和哈希值
	GetPublicKey(ctx context.Context) (key string, hash string, err error)
}

type PassportUseCase struct {
	repo PassportRepo
	log  *log.Helper
}

func NewPassportUseCase(repo PassportRepo, logger log.Logger) *PassportUseCase {
	return &PassportUseCase{repo: repo, log: log.NewHelper(logger)}
}

// GetKey 获取公钥和哈希
func (uc *PassportUseCase) GetKey(ctx context.Context) (key string, hash string, err error) {
	return uc.repo.GetPublicKey(ctx)
}

// CreatAccount 验证sid的md5值并创建用户签署登录token
func (uc *PassportUseCase) CreatAccount(ctx context.Context, sid string, key string) (id uint32, err error) {
	return uc.repo.CreatAccount(ctx, sid, key)
}

// PrepareCreateAccount 预创建账户
func (uc *PassportUseCase) PrepareCreateAccount(ctx context.Context, account Account, token string) error {

	// TODO 检测验证码token

	return uc.repo.PrepareCreateAccount(ctx, account)
}
