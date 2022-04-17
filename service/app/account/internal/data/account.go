package data

import (
	"context"
	v1 "service/api/account/v1"
	"service/app/account/internal/biz"
	"service/pkg/util/strconv"

	"github.com/go-kratos/kratos/v2/log"
)

var accountCacheKey = func(id uint32) string {
	return "account_model_" + strconv.UItoa(id)
}

var accountEMailCacheKey = func(email string) string {
	return "account_email_to_id_" + email
}

var accountUserIdKey = func(userid uint32) string {
	return "account_userid_to_id_" + strconv.UItoa(userid)
}

var modelToAccount = func(model *Account) *biz.Account {
	return &biz.Account{
		ID:           model.ID,
		UUID:         model.UUID,
		Email:        model.Email,
		Phone:        biz.TelPhone{TelCode: model.TelCode, Phone: model.Phone},
		PasswordHash: model.Password,
		Status:       model.Status,
		UserID:       model.UserID,
	}
}

type accountRepo struct {
	data *Data
	log  *log.Helper
}

// NewAccountRepo .
func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *accountRepo) CreateEMailAccount(ctx context.Context, account *biz.Account) (uint32, error) {
	model := &Account{
		UUID:     account.UUID,
		Email:    account.Email,
		TelCode:  account.Phone.TelCode,
		Phone:    account.Phone.Phone,
		Password: account.PasswordHash,
		Status:   account.Status,
		UserID:   0,
	}
	err := r.data.DataBase.Create(&model).Error
	if err != nil {
		return 0, err
	}
	err = r.SetAccountToCache(ctx, model)
	if err != nil {
		return 0, err
	}
	return model.ID, nil
}

func (r *accountRepo) GetAccountByUserID(ctx context.Context, userid uint32) (*biz.Account, error) {
	acc, ok, err := r.GetAccountByUserId(ctx, userid)
	if err != nil {
		return nil, err
	}
	if ok {
		return modelToAccount(acc), nil
	}

	acc, ok, err = r.GetAccountByUserIDFromDB(ctx, userid)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, v1.ErrorInternalServerError("用户不存在 %v", userid)
	}

	err = r.SetAccountToCache(ctx, acc)
	if err != nil {
		r.log.Error("缓存保存失败")
	}

	return modelToAccount(acc), nil
}

func (r *accountRepo) GetAccountByID(ctx context.Context, id uint32) (*biz.Account, error) {
	model, ok, err := r.GetAccountByIDFromCache(ctx, id)
	if err != nil {
		r.log.Error("") // TODO
	}
	if ok {
		return modelToAccount(model), nil
	}

	model, ok, err = r.GetAccountByIDFromDB(ctx, id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, v1.ErrorInternalServerError("用户不存在 %v", id) // TODO 用户不存在 err
	}

	err = r.SetAccountToCache(ctx, model)
	if err != nil {
		r.log.Error("") // TODO
	}

	return modelToAccount(model), nil
}

func (r *accountRepo) GetAccountByEMail(ctx context.Context, email string) (*biz.Account, error) {
	// 在缓存中查找
	model, ok, err := r.GetAccountByEMailFromCache(ctx, email)
	if err != nil {
		r.log.Error("") // TODO
	}
	if ok {
		return modelToAccount(model), nil
	}

	// 在数据库中查找
	model, ok, err = r.GetAccountByEMailFromDB(ctx, email)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, v1.ErrorInternalServerError("用户不存在 %v", email) // TODO 用户不存在 err
	}

	// 将账户模型保存到缓存中
	err = r.SetAccountToCache(ctx, model)
	if err != nil {
		r.log.Error("") // TODO
	}

	return modelToAccount(model), nil
}

func (r *accountRepo) GetAccountByPhone(ctx context.Context, phone *biz.TelPhone) (*biz.Account, error) {
	// TODO 未完成手机号功能
	return nil, nil
}

func (r *accountRepo) ExistAccountEMail(ctx context.Context, email string) (bool, error) {
	_, ok, err := r.data.Cache.GetString(ctx, accountEMailCacheKey(email))
	if err != nil {
		r.log.Error("") // TODO
	}
	if ok {
		return true, nil
	}

	// 通过邮箱查找账号数据
	model, ok, err := r.GetAccountByEMailFromDB(ctx, email)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	err = r.SetAccountToCache(ctx, model)
	if err != nil {
		r.log.Error("") // TODO
	}

	return true, nil
}

func (r *accountRepo) UpdateAccount(ctx context.Context, account *biz.Account) error {
	model, ok, err := r.GetAccountByIDFromDB(ctx, account.ID)
	if err != nil {
		return err
	}
	if !ok {
		return v1.ErrorInternalServerError("用户不存在 %v", account.ID) // TODO err
	}

	// 替换数据
	model.UUID = account.UUID
	model.TelCode = account.Phone.TelCode
	model.Phone = account.Phone.Phone
	model.Email = account.Email
	model.Status = account.Status
	model.Password = account.PasswordHash
	model.UserID = account.UserID
	// 储存到数据库并加入缓存
	err = r.data.DataBase.Save(&model).Error
	if err != nil {
		return err
	}
	err = r.DeleteAccountToCache(ctx, model)
	if err != nil {
		r.log.Error("") // TODO
	}
	return nil
}

func (r *accountRepo) BindUser(ctx context.Context, id, uid uint32) error {
	model, ok, err := r.GetAccountByIDFromDB(ctx, id)
	if err != nil {
		return err
	}
	if !ok {
		return v1.ErrorInternalServerError("用户不存在 %v", id) // TODO err
	}

	model.UserID = uid
	// 储存到数据库并加入缓存
	err = r.data.DataBase.Save(&model).Error
	if err != nil {
		return err
	}
	err = r.DeleteAccountToCache(ctx, model)
	if err != nil {
		r.log.Error("") // TODO
	}
	return nil
}
