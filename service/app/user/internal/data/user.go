package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"service/app/user/internal/biz"
	"service/pkg/util/strconv"
	"strings"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

var tagsToStringList = func(tags string) []string {
	return strings.Split(tags, ";")
}

var userCacheKey = func(id uint32) string {
	return "user_model_by_accountID_" + strconv.UItoa(id)
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
