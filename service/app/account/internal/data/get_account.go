package data

import (
	"context"
	v1 "service/api/account/v1"
	"service/app/account/internal/biz"
	"service/pkg/util/strconv"
)

func (r *accountRepo) GetAccountByID(ctx context.Context, id uint32) (*biz.Account, error) {
	var model Account
	ok, err := r.CacheGetAccountByID(ctx, &model, id)
	if err != nil {
		r.log.Error("") // TODO
	}
	if ok {
		return modelToAccount(&model), nil
	}

	ok, err = r.DBGetAccountByID(ctx, &model, id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, v1.ErrorInternalServerError("用户不存在 %v", id) // TODO 用户不存在 err
	}

	err = r.CacheSetAccount(ctx, &model)
	if err != nil {
		r.log.Error("") // TODO
	}

	return modelToAccount(&model), nil
}

func (r *accountRepo) GetAccountByEMail(ctx context.Context, email string) (*biz.Account, error) {
	var model Account
	// 在缓存中查找
	ok, err := r.CacheGetAccountByEMail(ctx, &model, email)
	if err != nil {
		r.log.Error("") // TODO
	}
	if ok {
		return modelToAccount(&model), nil
	}

	// 在数据库中查找
	ok, err = r.DBGetAccountByEMail(ctx, &model, email)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, v1.ErrorInternalServerError("用户不存在 %v", email) // TODO 用户不存在 err
	}

	// 将账户模型保存到缓存中
	err = r.CacheSetAccount(ctx, &model)
	if err != nil {
		r.log.Error("") // TODO
	}

	return modelToAccount(&model), nil
}

func (r *accountRepo) GetAccountByPhone(ctx context.Context, phone *biz.TelPhone) (*biz.Account, error) {
	// TODO 未完成手机号功能
	return nil, nil
}

// CacheGetAccountByID 使用Account主键ID从 缓存 中获取Account
func (r *accountRepo) CacheGetAccountByID(ctx context.Context, model *Account, id uint32) (ok bool, err error) {
	ok, err = r.data.Cache.Get(ctx, accountCacheKey(id), model)
	return
}

// CacheGetAccountByEMail 使用EMail从 缓存 中获取Account
func (r *accountRepo) CacheGetAccountByEMail(ctx context.Context, model *Account, email string) (ok bool, err error) {
	str, ok, err := r.data.Cache.GetString(ctx, accountEMailCacheKey(email))
	if err != nil {
		return
	}
	if !ok {
		return
	}

	id, err := strconv.ParseUint32(str)
	if err != nil {
		return false, err
	}
	return r.CacheGetAccountByID(ctx, model, id)
}

// DBGetAccountByID 使用Account主键ID从 数据库 中获取Account
func (r *accountRepo) DBGetAccountByID(ctx context.Context, model *Account, id uint32) (ok bool, err error) {
	tx := r.data.DataBase.Limit(1).Find(model, id)
	err = tx.Error
	if err != nil {
		return false, err
	}
	if tx.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

// DBGetAccountByEMail 使用EMail从 数据库 中获取Account
func (r *accountRepo) DBGetAccountByEMail(ctx context.Context, model *Account, email string) (ok bool, err error) {
	tx := r.data.DataBase.Limit(1).Find(model, "email = ?", email)
	err = tx.Error
	if err != nil {
		return false, err
	}
	if tx.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}
