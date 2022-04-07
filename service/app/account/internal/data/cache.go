package data

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"service/pkg/crypto/ecc"
	"service/pkg/util/strconv"
)

// GetAccountByUserId 通过用户id获取账户
func (r *accountRepo) GetAccountByUserId(
	ctx context.Context, id uint32) (model *Account, ok bool, err error) {

	str, ok, err := r.data.Cache.GetString(ctx, accountUserIdKey(id))
	if err != nil {
		return
	}

	if !ok {
		return
	}

	userid, err := strconv.ParseUint32(str)
	if err != nil {
		return
	}

	return r.GetAccountByUserIDFromDB(ctx, uint32(userid))
}

// GetAccountByIDFromCache 使用Account主键ID从缓存中获取Account
func (r *accountRepo) GetAccountByIDFromCache(ctx context.Context, id uint32) (model *Account, ok bool, err error) {
	ok, err = r.data.Cache.Get(ctx, accountCacheKey(id), &model)
	return
}

// GetAccountByEMailFromCache 使用EMail从缓存中获取Account
func (r *accountRepo) GetAccountByEMailFromCache(
	ctx context.Context, email string) (model *Account, ok bool, err error) {

	str, ok, err := r.data.Cache.GetString(ctx, accountEMailCacheKey(email))
	if err != nil {
		return
	}
	if !ok {
		return
	}

	id, err := strconv.ParseUint32(str)
	if err != nil {
		return nil, false, err
	}
	return r.GetAccountByIDFromCache(ctx, id)
}

// SetAccountToCache 保存账户数据到缓存
func (r *accountRepo) SetAccountToCache(ctx context.Context, model *Account) (err error) {
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

// DeleteAccountToCache 删除账户数据到缓存
func (r *accountRepo) DeleteAccountToCache(ctx context.Context, model *Account) (err error) {
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

// hashMd5To16 获取密钥md5 hash值，返回16个字符
var hashMd5To16 = func(privateKey string) string {
	bytes := md5.Sum([]byte(privateKey))
	return hex.EncodeToString(bytes[4:12])
}

// CreateLockOpenerToCache 创建钥匙对到缓存中
func (r *accountRepo) CreateLockOpenerToCache(ctx context.Context, id int) (lock *LockOpener, hash string, err error) {
	// 生成钥匙对
	privateKey, publicKey, err := ecc.GenerateKey()
	if err != nil {
		return
	}
	// 取密钥hash
	hash = hashMd5To16(privateKey)

	lock = &LockOpener{
		ID:      id,
		Public:  publicKey,
		Private: privateKey,
	}

	// 存缓存
	err = r.data.Cache.Set(ctx, lockOpenerCacheKey(hash), lock, 0)
	if err != nil {
		return
	}
	err = r.data.Cache.SetString(ctx, lockOpenerIDCacheKey(id), hash, 0)
	if err != nil {
		return
	}
	return
}
