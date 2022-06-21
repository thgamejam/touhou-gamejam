package data

import (
	"context"
	userV1 "service/api/user/v1"
	"service/app/user/internal/biz"
)

// GetUserByAccountID 通过账户ID获取用户信息
func (r *userRepo) GetUserByAccountID(ctx context.Context, accountID uint32) (user *biz.UserInfo, err error) {
	ok, userInfo, err := r.GetUserByAccountIDOnCache(ctx, accountID)
	if err != nil {
		return nil, err
	}
	if ok {
		return userInfo, nil
	}

	user, err = r.GetUserByAccountIDOnDatabase(ctx, accountID)
	return
}

// GetUserByAccountIDOnDatabase 通过账户ID在数据库中获取用户
func (r *userRepo) GetUserByAccountIDOnDatabase(ctx context.Context, accountID uint32) (user *biz.UserInfo, err error) {
	model := &User{}
	tx := r.data.DataBase.Find(model, "account_id = ?", accountID)
	if err != nil {
		return nil, err
	}
	if tx.RowsAffected == 0 {
		return nil, userV1.ErrorUserNotFoundByAccount("accountID : %v not found", accountID)
	}

	// 获取用户头像
	url, err := r.GetUserAvatarURL(ctx, model.AvatarID)
	if err != nil {
		r.log.Error("头像报错%v", err)
	}

	// TODO 标签ID转换为具体字段

	user = &biz.UserInfo{
		Name:      model.Name,
		AvatarUrl: url,
		WorkCount: model.WorksCount,
		FansCount: model.FansCount,
		Tags:      tagsToStringList(model.TagString),
	}

	err = r.data.Cache.Set(ctx, userCacheKey(accountID), user, 0)
	if err != nil {
		return nil, err
	}

	return
}

// GetUserByAccountIDOnCache 通过账户ID在缓存数据库中获取用户信息
func (r *userRepo) GetUserByAccountIDOnCache(ctx context.Context, accountID uint32) (ok bool, user *biz.UserInfo, err error) {
	ok, err = r.data.Cache.Get(ctx, userCacheKey(accountID), user)
	if err != nil {
		return false, nil, err
	}
	if !ok {
		return false, nil, nil
	}
	return
}

// GetUserAvatarURL 根据用户ID获取头像
func (r *userRepo) GetUserAvatarURL(ctx context.Context, avatarID string) (string, error) {
	// 从缓存中获取用户头像
	avatar, ok, _ := r.data.Cache.GetString(ctx, userAvatarIDCacheURL(avatarID))
	if ok {
		// 若缓存中存在则直接返回
		return avatar, nil
	}
	// 缓存中不存在则从对象存储里获取URL
	url, err := r.data.ObjectStorage.PreSignGetURL(ctx, r.conf.UserAvatarBucketName, avatarID, avatarID, -1)
	if err != nil {
		// 若对象存储报错则尝试获取默认头像
		url, err = r.data.ObjectStorage.PreSignGetURL(ctx, r.conf.UserAvatarBucketName, r.conf.DefaultUserAvatarHash, r.conf.DefaultUserAvatarHash, -1)
		return url.String(), err
	}
	// 添加用户头像url到缓存
	_ = r.data.Cache.SetString(ctx, userAvatarIDCacheURL(avatarID), url.String(), -1)

	return url.String(), err
}
