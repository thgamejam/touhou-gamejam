package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewAccountUseCase)

type AccountRepo interface {
	// CreateEMailAccount 创建邮箱账户
	// 返回创造账户的id
	CreateEMailAccount(context.Context, *Account) (uint64, error)

	// GetAccountByUserID 通过用户ID获取账户
	GetAccountByUserID(ctx context.Context, uint642 uint64) (*Account, error)

	// GetAccountByID 通过Account主键ID获取账户
	GetAccountByID(context.Context, uint64) (*Account, error)
	// GetAccountByEMail 通过用户邮箱获取账户
	GetAccountByEMail(context.Context, string) (*Account, error)
	// GetAccountByPhone 通过用户手机号获取账户
	GetAccountByPhone(context.Context, *TelPhone) (*Account, error)

	// ExistAccountEMail 是否存在邮箱
	ExistAccountEMail(context.Context, string) (bool, error)

	// GetPublicKey 获取公钥
	// 传入与公钥成对的密钥的md5-16哈希摘要
	GetPublicKey(context.Context, string) (*PublicKey, error)
	// GetRandomlyPublicKey 获取任意的一个公钥
	GetRandomlyPublicKey(context.Context) (*PublicKey, error)

	// GetPrivateKey 获取密钥
	// 传入密钥的md5-16哈希摘要
	GetPrivateKey(context.Context, string) (*PrivateKey, error)

	// UpdateAccount 更新账户
	UpdateAccount(context.Context, *Account) error

	// BindUser 绑定用户
	BindUser(context.Context, uint64, uint64) error
}

type AccountUseCase struct {
	repo AccountRepo
	log  *log.Helper
}

func NewAccountUseCase(repo AccountRepo, logger log.Logger) *AccountUseCase {
	return &AccountUseCase{repo: repo, log: log.NewHelper(logger)}
}
