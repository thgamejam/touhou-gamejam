package data

import "context"

func (r *accountRepo) ExistAccountEMail(ctx context.Context, email string) (bool, error) {
	_, ok, err := r.data.Cache.GetString(ctx, accountEMailCacheKey(email))
	if err != nil {
		r.log.Error("") // TODO
	}
	if ok {
		return true, nil
	}

	// 通过邮箱查找账号数据
	model, ok, err := r.DBGetAccountByEMail(ctx, email)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	err = r.CacheSetAccount(ctx, model)
	if err != nil {
		r.log.Error("") // TODO
	}

	return true, nil
}
