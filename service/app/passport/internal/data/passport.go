package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	accountV1 "service/api/account/v1"
	"service/app/passport/internal/biz"
)

type passportRepo struct {
	data *Data
	log  *log.Helper
}

// NewPassportRepo .
func NewPassportRepo(data *Data, logger log.Logger) biz.PassportRepo {
	return &passportRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *passportRepo) PrepareCreateAccount(ctx context.Context, account biz.Account) error {
	// grpc调用account服务预创建用户并返回预创建ID
	_, err := r.data.accountClient.PrepareCreateEMailAccount(ctx, &accountV1.PrepareCreateEMailAccountReq{
		Ciphertext: account.Password,
		Hash:       account.Hash,
		Email:      account.Email,
	})
	return err
}
