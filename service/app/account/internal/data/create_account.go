package data

import (
	"context"
	"service/app/account/internal/biz"
	"service/pkg/util/strconv"
)

func (r *accountRepo) CreateEMailAccount(ctx context.Context, account *biz.Account) (uint32, error) {
	model := &Account{
		UUID:     account.UUID,
		Email:    account.Email,
		TelCode:  account.Phone.TelCode,
		Phone:    account.Phone.Phone,
		Password: account.PasswordHash,
		Status:   account.Status,
		UserID:   0,
	}
	err := r.data.DataBase.Create(&model).Error
	if err != nil {
		return 0, err
	}
	err = r.CacheSetAccount(ctx, model)
	if err != nil {
		return 0, err
	}
	return model.ID, nil
}

// CacheSetAccount 保存账户数据到缓存
func (r *accountRepo) CacheSetAccount(ctx context.Context, model *Account) (err error) {
	// 维护邮箱和id关系 EMail: ID
	err = r.data.Cache.SetString(ctx, accountEMailCacheKey(model.Email), strconv.UItoa(model.ID), 0)
	if err != nil {
		return
	}

	// TODO 维护手机号和id的关系 未完成

	// 保存用户数据到缓存
	err = r.data.Cache.Set(ctx, accountCacheKey(model.ID), &model, 0)
	if err != nil {
		return
	}
	return
}
