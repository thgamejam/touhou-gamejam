package data

import (
    "context"
    "strconv"
)

// GetAccountByIDFromCache 使用ID从缓存中获取Account
func (r *accountRepo) GetAccountByIDFromCache(ctx context.Context, id uint64) (model *Account, ok bool, err error) {
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

    id, err := strconv.ParseUint(str, 10, 64)
    if err != nil {
        return nil, false, err
    }
    return r.GetAccountByIDFromCache(ctx, id)
}

// SetAccountToCache 保存账户数据到缓存
func (r *accountRepo) SetAccountToCache(ctx context.Context, model *Account) (err error) {
    // 维护邮箱和id关系 EMail: ID
    err = r.data.Cache.SetString(ctx, accountEMailCacheKey(model.Email), strconv.FormatUint(model.ID, 10), 0)
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
