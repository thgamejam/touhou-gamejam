package data

import (
	"context"
	v1 "service/api/account/v1"
	"service/app/account/internal/biz"
)

func (r *accountRepo) UpdateAccount(ctx context.Context, account *biz.Account) error {
	model, ok, err := r.DBGetAccountByID(ctx, account.ID)
	if err != nil {
		return err
	}
	if !ok {
		return v1.ErrorInternalServerError("用户不存在 %v", account.ID) // TODO err
	}

	// 替换数据
	model.UUID = account.UUID
	model.TelCode = account.Phone.TelCode
	model.Phone = account.Phone.Phone
	model.Email = account.Email
	model.Status = account.Status
	model.Password = account.PasswordHash
	// 储存到数据库并加入缓存
	err = r.data.DataBase.Save(&model).Error
	if err != nil {
		return err
	}
	err = r.CacheDeleteAccount(ctx, model)
	if err != nil {
		r.log.Error("") // TODO
	}
	return nil
}

// CacheDeleteAccount 删除账户数据到缓存
func (r *accountRepo) CacheDeleteAccount(ctx context.Context, model *Account) (err error) {
	// 删除邮箱和id关系 EMail: ID
	err = r.data.Cache.Del(ctx, accountEMailCacheKey(model.Email))
	if err != nil {
		return
	}

	// TODO 删除手机号和id的关系 未完成

	// 删除缓存中的用户数据
	err = r.data.Cache.Del(ctx, accountCacheKey(model.ID))
	if err != nil {
		return
	}
	return
}
