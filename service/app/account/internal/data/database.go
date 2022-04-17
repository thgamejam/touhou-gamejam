package data

import (
	"context"
)

func (r *accountRepo) GetAccountByUserIDFromDB(ctx context.Context, id uint32) (model *Account, ok bool, err error) {
	model = &Account{}
	tx := r.data.DataBase.Limit(1).Find(&model, "userid = ?", id)
	err = tx.Error
	if err != nil {
		return nil, false, err
	}
	if tx.RowsAffected == 0 {
		return nil, false, err
	}
	return model, true, nil
}

func (r *accountRepo) GetAccountByIDFromDB(ctx context.Context, id uint32) (model *Account, ok bool, err error) {
	model = &Account{}
	tx := r.data.DataBase.Limit(1).Find(&model, id)
	err = tx.Error
	if err != nil {
		return nil, false, err
	}
	if tx.RowsAffected == 0 {
		return nil, false, nil
	}
	return model, true, nil
}

func (r *accountRepo) GetAccountByEMailFromDB(ctx context.Context, email string) (model *Account, ok bool, err error) {
	model = &Account{}
	tx := r.data.DataBase.Limit(1).Find(&model, "email = ?", email)
	err = tx.Error
	if err != nil {
		return nil, false, err
	}
	if tx.RowsAffected == 0 {
		return nil, false, nil
	}
	return model, true, nil
}
