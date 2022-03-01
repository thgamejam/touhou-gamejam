package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "service/api/user/v1"
	"service/pkg/uuid"
	"strconv"
	"strings"
)

type User struct {
	ID       uint64
	Username string
	Password string
	*TelPhone
	Email string
	UUID  uuid.UUID // uuid 充当盐
}

// Account 账号
type Account struct {
	Format   v1.AccountFormat
	Username string
	Password string
}

type TelPhone struct {
	TelCode uint16
	Phone   string
}

type UserRepo interface {
	GetUserByUserID(context.Context, uint64) (*User, error)
	GetUserByUsername(context.Context, string) (*User, error)
	GetUserByPhone(context.Context, *TelPhone) (*User, error)
	GetUserByEmail(context.Context, string) (*User, error)

	VerifyPassword(context.Context, *User) (bool, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) error {
	return nil
}

// VerifyPassword 验证密码
// 注意:	用户名只允许"-"和"_"这两个符号
//		电话号码使用 国际区号+电话号 的格式, 86+18100000000
func (uc *UserUseCase) VerifyPassword(ctx context.Context, a *Account) (ok bool, err error) {
	var user *User
	switch a.Format {
	case v1.AccountFormat_ID:
		id, err := strconv.ParseUint(a.Username, 10, 64)
		if err != nil {
			return false, err
		}
		user, err = uc.repo.GetUserByUserID(ctx, id)
		if err != nil {
			return false, err
		}
		break

	case v1.AccountFormat_USERNAME:
		user, err = uc.repo.GetUserByUsername(ctx, a.Username)
		if err != nil {
			return false, err
		}
		break

	case v1.AccountFormat_EMAIL:
		user, err = uc.repo.GetUserByEmail(ctx, a.Username)
		if err != nil {
			return false, err
		}
		break

	case v1.AccountFormat_PHONE:
		// 拆分手机号, 86+18100000000
		phone := strings.Split(a.Username, "+")
		if len(phone) < 2 {
			return false, err
		}
		code, err := strconv.ParseUint(phone[0], 10, 16)
		if err != nil {
			return false, err
		}
		tel := &TelPhone{TelCode: uint16(code), Phone: phone[1]}
		user, err = uc.repo.GetUserByPhone(ctx, tel)
		if err != nil {
			return false, err
		}
		break

	default:
		return false, err
	}

	return uc.repo.VerifyPassword(ctx, user)
}
