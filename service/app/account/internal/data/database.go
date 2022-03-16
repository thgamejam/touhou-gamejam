package data

import (
    "context"
)

func (r *accountRepo) GetAccountByIDFromDB(ctx context.Context, id uint64) (model *Account, ok bool, err error) {
    tx := r.data.DataBase.First(&model, id)
    err = tx.Error
    ok = tx.RowsAffected == 0
    return
}

func (r *accountRepo) GetAccountByEMailFromDB(ctx context.Context, email string) (model *Account, ok bool, err error) {
    tx := r.data.DataBase.First(&model, "email = ?", email)
    err = tx.Error
    ok = tx.RowsAffected == 0
    return
}
