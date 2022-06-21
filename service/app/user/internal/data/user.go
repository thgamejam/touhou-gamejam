package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"service/app/user/internal/biz"
	"service/app/user/internal/conf"
	"service/pkg/util/strconv"
	"strings"
)

type userRepo struct {
	data *Data
	log  *log.Helper
	conf *conf.User
}

var tagsToStringList = func(tags string) []string {
	return strings.Split(tags, ";")
}

var userCacheKey = func(id uint32) string {
	return "user_model_by_accountID_" + strconv.UItoa(id)
}

var userAvatarIDCacheURL = func(hash string) string {
	return "user_AvatarID_URL_" + hash
}

// NewUserRepo .
func NewUserRepo(data *Data, conf *conf.User, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
		conf: conf,
	}
}
