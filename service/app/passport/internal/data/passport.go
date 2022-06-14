package data

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	accountV1 "service/api/account/v1"
	userV1 "service/api/user/v1"
	"service/app/passport/internal/biz"
	"service/app/passport/internal/conf"
	"service/pkg/jwt"
	"time"
)

type passportRepo struct {
	data *Data
	conf *conf.Passport
	log  *log.Helper
}

// NewPassportRepo .
func NewPassportRepo(data *Data, conf *conf.Passport, logger log.Logger) biz.PassportRepo {
	return &passportRepo{
		data: data,
		conf: conf,
		log:  log.NewHelper(logger),
	}
}

func (r *passportRepo) CreateUserByAccountID(ctx context.Context, id uint32) (ok bool, err error) {
	_, err = r.data.userClient.CreateUser(ctx, &userV1.CreateUserRequest{AccountID: id})
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetUserByAccountID 通过账户ID获取账户
func (r *passportRepo) GetUserByAccountID(ctx context.Context, id uint32) (ok bool, err error) {
	_, err = r.data.userClient.GetUserByAccountID(ctx, &userV1.GetUserByAccountIDRequest{AccountID: id})
	if err != nil {
		// 若错误信息为用户未找到则不返回错误
		if userV1.IsUserNotFoundByAccount(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// AccountLogout 注销会话号ID
func (r *passportRepo) AccountLogout(ctx context.Context, id uint32, sid string) (err error) {
	_, err = r.data.accountClient.CloseSession(ctx, &accountV1.CloseSessionReq{
		Id:  id,
		Sid: sid,
	})
	return
}

// ChangePassword 修改密码
func (r *passportRepo) ChangePassword(ctx context.Context, id uint32, ciphertext string, hash string) (err error) {
	_, err = r.data.accountClient.SavePassword(ctx, &accountV1.SavePasswordReq{
		Id:         id,
		Ciphertext: ciphertext,
		Hash:       hash,
	})
	if err != nil {
		return err
	}
	return nil
}

// VerifyAccountTokenId 验证会话Token是否合法
func (r *passportRepo) VerifyAccountTokenId(ctx context.Context) (accountId uint32, err error) {
	// 格式化ctx
	tr, y := transport.FromServerContext(ctx)
	if !y {
		// 错误 格式化错误
		return 0, errors.New("FromServerError")
	}
	// 获取token字符
	token := tr.RequestHeader().Get("Authorization")
	if token == "" {
		// 错误 Token没有找到
		return 0, errors.New("TokenNotFound")
	}
	// 解密Token
	loginToken, success := jwt.ValidateLoginToken(token, []byte(r.conf.VerifyEmailKey))
	if !success {
		// 错误 Token解析错误
		return 0, errors.New("TokenValidateError")
	}
	if loginToken.RenewalAt < time.Now().Unix() {
		// 错误 时间过期
		return 0, errors.New("PleaseRenewalToken")
	}
	// 验证会话ID是否合法
	session, err := r.data.accountClient.VerifySession(ctx, &accountV1.VerifySessionReq{
		Id:  loginToken.UserID,
		Sid: loginToken.UUID,
	})
	if err != nil {
		return 0, err
	}
	if !session.Ok {
		// 错误 会话ID不合法
		return 0, errors.New("PleaseRenewalToken")
	}
	// 返回合法的用户ID
	return loginToken.UserID, nil
}

// ChangeUserPassword 修改用户密码
func (r *passportRepo) ChangeUserPassword(ctx context.Context, accountID uint32, password string, hash string) (ok bool, err error) {
	_, err = r.data.accountClient.SavePassword(ctx, &accountV1.SavePasswordReq{
		Id:         accountID,
		Ciphertext: password,
		Hash:       hash,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// LoginVerify 登录验证
func (r *passportRepo) LoginVerify(ctx context.Context, username string, ciphertext string, hash string) (uint32, error) {
	userID, err := r.data.accountClient.VerifyPassword(ctx, &accountV1.VerifyPasswordReq{
		Username:   username,
		Ciphertext: ciphertext,
		Hash:       hash,
	})
	if err != nil {
		return 0, err
	}

	return userID.Id, nil
}

// SignLoginToken 签署登录token
func (r *passportRepo) SignLoginToken(ctx context.Context, accountID uint32) (token string, err error) {
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return "", errors.New("ctxError")
	}
	ip := tr.RequestHeader().Get("X-RemoteAddr")
	sid, err := r.data.accountClient.CreateSession(ctx, &accountV1.CreateSessionReq{
		Id:        accountID,
		Ip:        ip,
		ExpiresAt: r.conf.LoginExpireTime,
	})
	if err != nil {
		return "", err
	}
	t, err := jwt.CreateLoginToken(jwt.LoginToken{
		UserID:    accountID,
		UUID:      sid.Sid,
		CreateAt:  time.Now().Unix(),
		RenewalAt: time.Now().Add(r.conf.RenewalTime.AsDuration()).Unix(),
	}, []byte(r.conf.VerifyEmailKey), r.conf.RenewalTime)
	if err != nil {
		return "", err
	}
	return t, nil
}

// GetPublicKey 获取公钥和哈希值
func (r *passportRepo) GetPublicKey(ctx context.Context) (key string, hash string, err error) {
	rep, err := r.data.accountClient.GetKey(ctx, &accountV1.GetKeyReq{})
	if err != nil {
		return
	}
	return rep.Key, rep.Hash, nil
}

// CreatAccount 创建用户
func (r *passportRepo) CreatAccount(ctx context.Context, sid string, key string) (uint32, error) {
	sidMd5 := md5.Sum([]byte(sid + r.conf.VerifyEmailKey))
	keyHash := hex.EncodeToString(sidMd5[:])
	if key != keyHash {
		return 0, errors.New("") // TODO Creat Account ERROR
	}

	accountID, err := r.data.accountClient.FinishCreateEMailAccount(ctx, &accountV1.FinishCreateEMailAccountReq{Sid: sid})
	if err != nil {
		return 0, err
	}
	return accountID.Id, nil
}

// PrepareCreateAccount 预创建账户
func (r *passportRepo) PrepareCreateAccount(ctx context.Context, account biz.Account) error {
	// grpc调用account服务预创建用户并返回预创建ID
	_, err := r.data.accountClient.PrepareCreateEMailAccount(ctx, &accountV1.PrepareCreateEMailAccountReq{
		Ciphertext: account.Password,
		Hash:       account.Hash,
		Email:      account.Email,
	})
	return err
}
