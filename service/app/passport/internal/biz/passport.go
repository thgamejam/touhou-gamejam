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
	// LoginVerify 登录校验
	LoginVerify(ctx context.Context, username string, ciphertext string, hash string) (id uint32, err error)
	// ChangeUserPassword 修改用户密码
	ChangeUserPassword(ctx context.Context, accountID uint32, password string, hash string) (ok bool, err error)
	// VerifyAccountTokenId 验证用户Token是否合法并返回合法账户ID
	VerifyAccountTokenId(ctx context.Context) (accountId uint32, err error)
	// ChangePassword 修改密码
	ChangePassword(ctx context.Context, id uint32, ciphertext string, hash string) (err error)
	// AccountLogout 注销会话号ID
	AccountLogout(ctx context.Context, id uint32, sid string) (err error)
}

type PassportUseCase struct {
	repo PassportRepo
	log  *log.Helper
}

// Logout 登出
func (uc *PassportUseCase) Logout(ctx context.Context, id uint32, sid string) (err error) {
	err = uc.repo.AccountLogout(ctx, id, sid)
	return
}

// ChangePassword 修改密码
func (uc *PassportUseCase) ChangePassword(ctx context.Context, id uint32, ciphertext string, hash string) (token string, err error) {
	err = uc.repo.ChangePassword(ctx, id, ciphertext, hash)
	if err != nil {
		return "", err
	}
	token, err = uc.repo.SignLoginToken(ctx, id)
	if err != nil {
		return "", err
	}
	return
}

func NewPassportUseCase(repo PassportRepo, logger log.Logger) *PassportUseCase {
	return &PassportUseCase{repo: repo, log: log.NewHelper(logger)}
}

// RenewalToken 验证并续签Token
func (uc *PassportUseCase) RenewalToken(ctx context.Context) (string, error) {
	accountId, err := uc.repo.VerifyAccountTokenId(ctx)
	if err != nil {
		return "", err
	}
	token, err := uc.repo.SignLoginToken(ctx, accountId)
	if err != nil {
		return "", err
	}
	return token, nil
}

// GetKey 获取公钥和哈希
func (uc *PassportUseCase) GetKey(ctx context.Context) (key string, hash string, err error) {
	return uc.repo.GetPublicKey(ctx)
}

// CreatAccount 验证sid的md5值并创建用户签署登录token
func (uc *PassportUseCase) CreatAccount(ctx context.Context, sid string, key string) (token string, err error) {
	id, err := uc.repo.CreatAccount(ctx, sid, key)
	if err != nil {
		return "", err
	}
	token, err = uc.repo.SignLoginToken(ctx, id)
	return
}

// Login 登录并签署验证token
func (uc *PassportUseCase) Login(ctx context.Context, username string, ciphertext string, hash string) (ok bool, token string, err error) {
	userID, err := uc.repo.LoginVerify(ctx, username, ciphertext, hash)
	if err != nil {
		return false, "", err
	}

	// TODO 判断user是否存在，若不存在则新建user

	loginToken, err := uc.repo.SignLoginToken(ctx, userID)
	if err != nil {
		return false, "", err
	}

	return true, loginToken, nil
}

// PrepareCreateAccount 预创建账户
func (uc *PassportUseCase) PrepareCreateAccount(ctx context.Context, account Account, token string) error {
	// TODO 验证人机token
	return uc.repo.PrepareCreateAccount(ctx, account)
}
