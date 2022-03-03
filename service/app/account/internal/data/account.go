package data

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "service/app/account/internal/biz"
    "strconv"
)

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

var accountCacheKey = func(id uint64) string {
    return "account_model_" + strconv.FormatUint(id, 10)
}

var accountEMailCacheKey = func(email string) string {
    return "account_email_to_id_" + email
}

var modelToAccount = func(a *Account) *biz.Account {
    return &biz.Account{
        ID:           a.ID,
        UUID:         a.UUID,
        Email:        a.Email,
        Phone:        &biz.TelPhone{TelCode: a.TelCode, Phone: a.Phone},
        PasswordHash: a.Password,
        Status:       a.Status,
    }
}

func (r *accountRepo) CreateEMailAccount(ctx context.Context, account *biz.Account) error {
    model := &Account{
        UUID:     account.UUID,
        Email:    account.Email,
        TelCode:  account.Phone.TelCode,
        Phone:    account.Phone.Phone,
        Password: account.PasswordHash,
        Status:   account.Status,
    }
    return r.data.DataBase.Save(&model).Error
}

// saveAccountModelToCache 保存账户数据到缓存
func (r *accountRepo) saveAccountModelToCache(ctx context.Context, account *Account) (err error) {
    // 维护邮箱和id关系 EMail: ID
    err = r.data.Cache.SaveString(ctx, accountEMailCacheKey(account.Email), strconv.FormatUint(account.ID, 10), 0)
    if err != nil {
        return
    }
    
    // TODO 维护手机号和id的关系 未完成
    
    // 保存用户数据到缓存
    err = r.data.Cache.Save(ctx, accountCacheKey(account.ID), &account, 0)
    if err != nil {
        return
    }
    return
}

func (r *accountRepo) GetAccountByID(ctx context.Context, id uint64) (*biz.Account, error) {
    account := &Account{}
    ok, err := r.data.Cache.Get(ctx, accountCacheKey(id), &account)
    if err != nil {
        return nil, err
    }
    // 缓存不存在key
    if !ok {
        err = r.data.DataBase.First(&account, id).Error
        if err == nil {
            return nil, err
        }
        err = r.saveAccountModelToCache(ctx, account)
        if err != nil {
            return nil, err
        }
    }
    return modelToAccount(account), nil
}

func (r *accountRepo) GetAccountByEMail(ctx context.Context, email string) (*biz.Account, error) {
    account := &Account{}
    v , ok, err := r.data.Cache.GetString(ctx, accountEMailCacheKey(email))
    if err != nil {
        log.Error("") // TODO
    }
    
    if ok {
        var id, err = strconv.ParseUint(v, 10, 64)
        if err != nil {
            return nil, err
        }
        return r.GetAccountByID(ctx, id)
    }
    
    err = r.data.DataBase.First(&account, "email = ?", email).Error
    if err != nil {
        return  nil, err
    }
    
    err = r.saveAccountModelToCache(ctx, account)
    if err != nil {
        log.Error("")  // TODO
    }
    
    return modelToAccount(account), nil
}

func (r *accountRepo) GetAccountByPhone(ctx context.Context, phone *biz.TelPhone) (*biz.Account, error) {
    // TODO 未完成手机号功能
    return nil, nil
}

func (r *accountRepo) GetPublicKey(ctx context.Context, s string) (*biz.PublicKey, error) {
    panic("implement me")
}

func (r *accountRepo) GetPrivateKey(ctx context.Context, s string) (*biz.PrivateKey, error) {
    panic("implement me")
}

func (r *accountRepo) UpdateAccount(ctx context.Context, account *biz.Account) error {
    if _, err := r.GetAccountByID(ctx, account.ID); err != nil {
        return err
    }
    err := r.data.DataBase.Save(&account).Error
    if err != nil {
        return err
    }
    err = r.data.Cache.Save(ctx, accountCacheKey(account.ID), &account, 0)
    return err
}
