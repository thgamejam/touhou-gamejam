package data

import (
	"context"
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
		return nil, nil
	}

	// TODO 将头像ID转换为URL 标签ID转换为具体字段

	user = &biz.UserInfo{
		Name:      model.Name,
		AvatarUrl: "model.AvatarID",
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
