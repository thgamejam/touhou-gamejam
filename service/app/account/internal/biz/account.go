package biz

import (
	"bytes"
	"context"
	"crypto/sha1"
	v1 "service/api/account/v1"
	"service/pkg/crypto/ecc"
	"service/pkg/uuid"
)

// PasswordCiphertext 密码密文, 被公钥加密的密码
type PasswordCiphertext struct {
	KeyHash    string // 密钥摘要
	Ciphertext string // 密码密文
}

// Account 账户
type Account struct {
	ID           uint32
	UUID         uuid.UUID // 唯一标识符
	Email        string    // 邮箱
	Phone        TelPhone  // 电话号码
	PasswordHash []byte    // 密码哈希值
	Status       uint8     // 状态
}

// TelPhone 电话号码
type TelPhone struct {
	TelCode uint16 // 国际区号
	Phone   string // 电话号码
}

// decryptPassword 密码解码器
var decryptPassword = func(key *PrivateKey, ciphertext string) (plaintext string, err error) {
	ecdsaKey, err := ecc.ParsePrivateKey(key.Key)
	if err != nil {
		return
	}
	bytes, err := ecc.Decrypt(ecdsaKey, ciphertext)
	if err != nil {
		return
	}
	return string(bytes), nil
}

// hashPassword 密码取哈希值
var hashPassword = func(password, sign string) []byte {
	byte20 := sha1.Sum([]byte(password + sign))
	return byte20[:]
}

// PrepareCreateEMailAccount 预创建用户
func (uc *AccountUseCase) PrepareCreateEMailAccount(
	ctx context.Context, email string, passwdCT *PasswordCiphertext) (sid string, err error) {

	// 验证邮箱是否存在
	isExist, err := uc.repo.ExistAccountEMail(ctx, email)
	if err != nil {
		return "", err
	}
	if isExist {
		return "", v1.ErrorEmailAlreadyExists("邮箱已注册 (%v)", email)
	}

	return uc.repo.SavePrepareCreateEMailAccount(ctx, email, passwdCT)
}

// FinishCreateEMailAccount 完成邮箱创建账户
func (uc *AccountUseCase) FinishCreateEMailAccount(ctx context.Context, sid string) (id uint32, err error) {

	email, ciphertext, err := uc.repo.GetAndDeletePrepareCreateEMailAccount(ctx, sid)
	if err != nil {
		return 0, err
	}

	isExist, err := uc.repo.ExistAccountEMail(ctx, email)
	if err != nil {
		return 0, err
	}
	// 防止账户被创建多次
	if isExist {
		return 0, v1.ErrorEmailAlreadyExists("邮箱已注册 (%v)", email)
	}

	// 解密
	password, err := uc.getPasswordPlaintext(ctx, ciphertext)
	if err != nil {
		return
	}
	// 创建有序的唯一id
	guid, err := uuid.NewOrderedUUID()
	if err != nil {
		return 0, err
	}
	// 取密码哈希
	passwordHash := hashPassword(password, guid.String())

	a := &Account{
		UUID:         guid,
		Email:        email,
		PasswordHash: passwordHash,
		Status:       0,
	}

	return uc.repo.CreateEMailAccount(ctx, a)
}

// getPasswordPlaintext 获取密码明文
func (uc *AccountUseCase) getPasswordPlaintext(ctx context.Context, passwdCT *PasswordCiphertext) (string, error) {
	// 获取私钥
	key, err := uc.repo.GetPrivateKey(ctx, passwdCT.KeyHash)
	if err != nil {
		return "", err
	}
	// 解密
	return decryptPassword(key, passwdCT.Ciphertext)
}

// GetAccount 通过ID获取账号
func (uc *AccountUseCase) GetAccount(ctx context.Context, id uint32) (*Account, error) {
	return uc.repo.GetAccountByID(ctx, id)
}

// SavePassword 保存密码, 修改密码
func (uc *AccountUseCase) SavePassword(ctx context.Context, id uint32, passwdCT *PasswordCiphertext) (err error) {
	// 获取账户
	account, err := uc.repo.GetAccountByID(ctx, id)
	if err != nil {
		return
	}
	// 解密
	password, err := uc.getPasswordPlaintext(ctx, passwdCT)
	if err != nil {
		return
	}
	// 更新uuid
	account.UUID, err = uuid.NewOrderedUUID()
	if err != nil {
		return
	}
	// 取密码哈希
	hash := hashPassword(password, account.UUID.String())
	// 更新
	account.PasswordHash = hash
	err = uc.repo.UpdateAccount(ctx, account)
	return
}

// VerifyPasswordByEMail 通过邮箱验证对应账户的密码
func (uc *AccountUseCase) VerifyPasswordByEMail(
	ctx context.Context, email string, passwdCT *PasswordCiphertext) (id uint32, ok bool, err error) {
	account, err := uc.repo.GetAccountByEMail(ctx, email)
	if err != nil {
		return
	}
	// 解密
	password, err := uc.getPasswordPlaintext(ctx, passwdCT)
	if err != nil {
		return
	}
	// 取密码哈希
	hash := hashPassword(password, account.UUID.String())
	ok = bytes.Equal(account.PasswordHash, hash)
	id = account.ID
	return
}

// ExistAccountEMail 是否存在邮箱
func (uc *AccountUseCase) ExistAccountEMail(ctx context.Context, email string) (bool, error) {
	return uc.repo.ExistAccountEMail(ctx, email)
}
