package data

import (
	"context"
	"service/app/user/internal/biz"
)

// CreateUser 创建用户
func (r *userRepo) CreateUser(ctx context.Context, accountID uint32) (user *biz.UserInfo, err error) {

	// TODO 初始化用户信息

	model := User{
		Name:             "",
		AccountID:        accountID,
		AvatarID:         "",
		TagString:        "",
		AllowSyndication: true,
		WorksCount:       0,
		FansCount:        0,
	}
	err = r.data.DataBase.Create(&model).Error
	if err != nil {
		return nil, err
	}
	user = &biz.UserInfo{
		Name:      model.Name,
		AvatarUrl: model.AvatarID,
		WorkCount: model.WorksCount,
		FansCount: model.FansCount,
		Tags:      tagsToStringList(model.TagString),
	}
	_ = r.data.Cache.Set(ctx, userCacheKey(accountID), *user, 0)
	return
}
