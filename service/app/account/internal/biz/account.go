package biz

import (
    "context"
    "crypto/sha1"
    "encoding/hex"
    "github.com/go-kratos/kratos/v2/log"
    "service/pkg/crypto/ecc"
    "service/pkg/uuid"
)

// PublicKey 公钥
type PublicKey struct {
    Hash string // 密钥摘要
    Key  string // 公钥内容
}

// PrivateKey 密钥
type PrivateKey struct {
    Hash string // 密钥摘要
    Key  string // 密钥内容
}

// PasswordCiphertext 密码密文, 被公钥加密的密码
type PasswordCiphertext struct {
    KeyHash    string // 密钥摘要
    Ciphertext string // 密码密文
}

// Account 账户
type Account struct {
    ID           uint64
    UUID         uuid.UUID // 唯一标识符
    Email        string    // 邮箱
    Phone        *TelPhone // 电话号码
    PasswordHash string    // 密码哈希值
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
var hashPassword = func(password, sign string) string {
    hashByte := sha1.Sum([]byte(password + sign))
    return hex.EncodeToString(hashByte[:])
}

type AccountRepo interface {
    // CreateEMailAccount 创建邮箱账户
    // 返回创造账户的id
    CreateEMailAccount(context.Context, *Account) (uint64, error)

    // GetAccountByID 通过用户ID获取账户
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
}

type AccountUseCase struct {
    repo AccountRepo
    log  *log.Helper
}

func NewAccountUseCase(repo AccountRepo, logger log.Logger) *AccountUseCase {
    return &AccountUseCase{repo: repo, log: log.NewHelper(logger)}
}

// GetKey 获取公钥
func (uc *AccountUseCase) GetKey(ctx context.Context, hash string) (*PublicKey, error) {
    return uc.repo.GetPublicKey(ctx, hash)
}

// GetRandomlyKey 获取任意的一个公钥
func (uc *AccountUseCase) GetRandomlyKey(ctx context.Context) (*PublicKey, error) {
    return uc.repo.GetRandomlyPublicKey(ctx)
}

// CreateEMailAccount 使用邮箱创建账户
func (uc *AccountUseCase) CreateEMailAccount(
    ctx context.Context, email string, passwdCT *PasswordCiphertext) (id uint64, err error) {

    // TODO 未判断邮箱是否已被注册

    // 解密
    password, err := uc.getPasswordPlaintext(ctx, passwdCT)
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
func (uc *AccountUseCase) GetAccount(ctx context.Context, id uint64) (*Account, error) {
    return uc.repo.GetAccountByID(ctx, id)
}

// SavePassword 保存密码, 修改密码
func (uc *AccountUseCase) SavePassword(ctx context.Context, id uint64, passwdCT *PasswordCiphertext) (err error) {
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
    ctx context.Context, email string, passwdCT *PasswordCiphertext) (id uint64, ok bool, err error) {
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
    ok = account.PasswordHash == hash
    id = account.ID
    return
}

// ExistAccountEMail 是否存在邮箱
func (uc *AccountUseCase) ExistAccountEMail(ctx context.Context, email string) (bool, error) {
    return uc.ExistAccountEMail(ctx, email)
}
