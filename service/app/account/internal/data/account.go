package data

import (
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
