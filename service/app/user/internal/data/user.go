package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"service/app/user/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

var userCacheKey = func(username string) string {
	return "user_" + username
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
