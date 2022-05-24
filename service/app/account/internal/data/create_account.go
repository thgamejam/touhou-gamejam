package data

import (
	"context"
	"errors"
	"service/app/account/internal/biz"
	"service/pkg/util/strconv"
	uuid2 "service/pkg/uuid"
)

var prepareCreateEMailAccountCacheKey = func(sid string) string {
	return "prepare_create_account_to_uid_" + sid
}

// SavePrepareCreateEMailAccount 保存预创建账户数据
func (r *accountRepo) SavePrepareCreateEMailAccount(
	ctx context.Context, email string, ciphertext *biz.PasswordCiphertext) (sid string, err error) {
	// 生成会话号
	sid = uuid2.New().String()

	cache := &PrepareCreateEMailAccountCache{
		Email:      email,
		KeyHash:    ciphertext.KeyHash,
		Ciphertext: ciphertext.Ciphertext,
	}

	// 保存预创建账户到缓存
	err = r.data.Cache.Set(ctx, prepareCreateEMailAccountCacheKey(sid), &cache, 0)
	if err != nil {
		return "", err
	}

	return sid, err
}

// GetAndDeletePrepareCreateEMailAccount 获取并删除保存的预创建账户数据
func (r *accountRepo) GetAndDeletePrepareCreateEMailAccount(
	ctx context.Context, sid string) (email string, ciphertext *biz.PasswordCiphertext, err error) {

	var cache *PrepareCreateEMailAccountCache
	ok, err := r.data.Cache.Get(ctx, prepareCreateEMailAccountCacheKey(sid), cache)
	if err != nil {
		return "", nil, err
	}
	if !ok {
		return "", nil, errors.New("") // TODO Get Prepare Create EMail Account ERROR
	}

	r.data.Cache.Del(ctx, prepareCreateEMailAccountCacheKey(sid))

	return cache.Email, &biz.PasswordCiphertext{KeyHash: cache.KeyHash, Ciphertext: cache.Ciphertext}, nil
}

// CreateEMailAccount 创建邮箱账户
func (r *accountRepo) CreateEMailAccount(ctx context.Context, account *biz.Account) (uint32, error) {
	model := &Account{
		UUID:     account.UUID,
		Email:    account.Email,
		TelCode:  account.Phone.TelCode,
		Phone:    account.Phone.Phone,
		Password: account.PasswordHash,
		Status:   account.Status,
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
