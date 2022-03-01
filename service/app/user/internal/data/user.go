package data

import (
    "context"
    "encoding/json"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-redis/redis/v8"
    "service/app/user/internal/biz"
    "service/pkg/uuid"
)

var _ biz.UserRepo = (*userRepo)(nil)

var userCacheKey = func(uuid uuid.UUID) string {
    return "user_" + uuid.String()
}

type userRepo struct {
    data *Data
    log  *log.Helper
}

func (u userRepo) GetUserByUserID(ctx context.Context, u2 uint64) (*biz.User, error) {
    panic("implement me")
}

func (u userRepo) GetUserByUserUUID(ctx context.Context, uuid uuid.UUID) (*biz.User, error) {
    account := &Account{
        UUID: uuid,
    }
    accountInfo, err := u.data.Redis.Get(ctx, userCacheKey(uuid)).Result()
    if err == redis.Nil {
        // 缓存不存在，从数据库获取数据
        if err = u.data.DataBase.First(&account).Error; err != nil {
            return nil, err
        }
        // 序列化数据
        accountByte, err := json.Marshal(account)
        if err == nil {
            return nil, err
        }
        accountInfo = string(accountByte)
        // 存入redis缓存
        err = u.data.Redis.Set(ctx, userCacheKey(uuid), accountInfo, 0).Err()
        if err == nil {
            return nil, err
        }
    } else {
        // 缓存存在
        err = json.Unmarshal([]byte(accountInfo), &account)
    }
    
    return &biz.User{
        ID:       account.ID,
        Username: account.Username,
        Password: account.Password,
        Email:    account.Email,
        UUID:     account.UUID,
		TelPhone: &biz.TelPhone{
			TelCode: account.TelCode,
			Phone: account.Phone,
		},
    }, err
}

func (u userRepo) GetUserByUsername(ctx context.Context, s string) (*biz.User, error) {
    panic("implement me")
}

func (u userRepo) GetUserByPhone(ctx context.Context, phone *biz.TelPhone) (*biz.User, error) {
    panic("implement me")
}

func (u userRepo) GetUserByEmail(ctx context.Context, s string) (*biz.User, error) {
    panic("implement me")
}

func (u userRepo) VerifyPassword(ctx context.Context, user *biz.User) (bool, error) {
    panic("implement me")
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
    return &userRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}
