package data

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	accountV1 "service/api/account/v1"
	"service/app/passport/internal/biz"
	"service/app/passport/internal/conf"
	"service/pkg/jwt"
	"time"
)

type passportRepo struct {
	data *Data
	conf *conf.Passport
	log  *log.Helper
}

// NewPassportRepo .
func NewPassportRepo(data *Data, conf *conf.Passport, logger log.Logger) biz.PassportRepo {
	return &passportRepo{
		data: data,
		conf: conf,
		log:  log.NewHelper(logger),
	}
}

// SignLoginToken 签署登录token
func (r *passportRepo) SignLoginToken(ctx context.Context, accountID uint32) (token string, err error) {
	t, err := jwt.CreateLoginToken(jwt.LoginToken{
		UserID:     accountID,
		CreateTime: time.Now().Unix(),
	}, []byte(r.conf.VerifyEmailKey), time.Duration(r.conf.LoginExpireTime)*time.Second)
	if err != nil {
		return "", err
	}
	return t, nil
}

// GetPublicKey 获取公钥和哈希值
func (r *passportRepo) GetPublicKey(ctx context.Context) (key string, hash string, err error) {
	rep, err := r.data.accountClient.GetKey(ctx, &accountV1.GetKeyReq{})
	if err != nil {
		return
	}
	return rep.Key, rep.Hash, nil
}

// CreatAccount 创建用户
func (r *passportRepo) CreatAccount(ctx context.Context, sid string, key string) (uint32, error) {
	sidMd5 := md5.Sum([]byte(sid + r.conf.VerifyEmailKey))
	keyHash := hex.EncodeToString(sidMd5[:])
	if key != keyHash {
		return 0, errors.New("") // TODO Creat Account ERROR
	}

	accountID, err := r.data.accountClient.FinishCreateEMailAccount(ctx, &accountV1.FinishCreateEMailAccountReq{Sid: sid})
	if err != nil {
		return 0, err
	}
	return accountID.Id, nil
}

// PrepareCreateAccount 预创建账户
func (r *passportRepo) PrepareCreateAccount(ctx context.Context, account biz.Account) error {
	// grpc调用account服务预创建用户并返回预创建ID
	_, err := r.data.accountClient.PrepareCreateEMailAccount(ctx, &accountV1.PrepareCreateEMailAccountReq{
		Ciphertext: account.Password,
		Hash:       account.Hash,
		Email:      account.Email,
	})
	return err
}
