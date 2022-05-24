package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewAccountUseCase)

type AccountRepo interface {
	// SavePrepareCreateEMailAccount 保存预创建账户数据
	// 返回一个预创建的会话号
	SavePrepareCreateEMailAccount(ctx context.Context, email string, ciphertext *PasswordCiphertext) (sid string, err error)

	// GetPrepareCreateEMailAccount 获取保存的预创建账户数据
	// 需要保存时的会话号
	GetPrepareCreateEMailAccount(ctx context.Context, sid string) (email string, ciphertext *PasswordCiphertext, err error)

	// CreateEMailAccount 创建邮箱账户
	// 返回创造账户的id
	CreateEMailAccount(ctx context.Context, account *Account) (id uint32, err error)

	// GetAccountByID 通过Account主键ID获取账户
	GetAccountByID(ctx context.Context, id uint32) (*Account, error)
	// GetAccountByEMail 通过用户邮箱获取账户
	GetAccountByEMail(ctx context.Context, email string) (*Account, error)
	// GetAccountByPhone 通过用户手机号获取账户
	GetAccountByPhone(ctx context.Context, phone *TelPhone) (*Account, error)

	// ExistAccountEMail 是否存在邮箱
	ExistAccountEMail(ctx context.Context, email string) (ok bool, err error)

	// GetPublicKey 获取公钥
	// 传入与公钥成对的密钥的md5-16哈希摘要
	GetPublicKey(ctx context.Context, hash string) (*PublicKey, error)
	// GetRandomlyPublicKey 获取任意的一个公钥
	GetRandomlyPublicKey(ctx context.Context) (*PublicKey, error)

	// GetPrivateKey 获取密钥
	// 传入密钥的md5-16哈希摘要
	GetPrivateKey(ctx context.Context, hash string) (*PrivateKey, error)

	// UpdateAccount 更新账户
	UpdateAccount(ctx context.Context, account *Account) error
}

type AccountUseCase struct {
	repo AccountRepo
	log  *log.Helper
}

func NewAccountUseCase(repo AccountRepo, logger log.Logger) *AccountUseCase {
	return &AccountUseCase{repo: repo, log: log.NewHelper(logger)}
}
