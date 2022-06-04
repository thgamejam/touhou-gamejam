package jwt

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/protobuf/ptypes/duration"
	"time"
)

type (
	// loginTokenKey context中登录token的键
	loginTokenKey struct{}
)

// LoginToken 登录token
type LoginToken struct {
	UserID    uint32 `json:"user_id"`    // 用户id
	UUID      string `json:"uuid"`       // 会话号
	CreateAt  int64  `json:"create_At"`  // 创建时间
	RenewalAt int64  `json:"renewal_At"` // 续签到期时间, 保活时间
	jwt.RegisteredClaims
}

// FromLoginTokenContext 返回存储在context中的登录token，如果有
func FromLoginTokenContext(ctx context.Context) (loginToken *LoginToken, ok bool) {
	loginToken, ok = ctx.Value(loginTokenKey{}).(*LoginToken)
	return
}

// ValidateLoginListMatcher 登录鉴权路由匹配器
func ValidateLoginListMatcher(routers []string) selector.MatchFunc {
	authRouters := make(map[string]bool)
	for i := 0; i < len(routers); i++ {
		authRouters[routers[i]] = true
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := authRouters[operation]; ok {
			return true
		}
		return false
	}
}

// LoginAuthMiddleware 登录鉴权中间件
func LoginAuthMiddleware(secret []byte) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				token := tr.RequestHeader().Get("Authorization")
				if token == "" {
					return nil, errors.New("TokenNotFound")
				}
				loginToken, success := ValidateLoginToken(token, secret)
				if !success {
					return nil, errors.New("TokenValidateError")
				}
				if loginToken.RenewalAt < time.Now().Unix() {
					return nil, errors.New("PleaseRenewalToken")
				}
				ctx = context.WithValue(ctx, loginTokenKey{}, loginToken)
			}

			return handler(ctx, req)
		}
	}
}

// CreateLoginToken 创建登录token
func CreateLoginToken(claims LoginToken, secret []byte, expirationTime *duration.Duration) (signedToken string, err error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(expirationTime.AsDuration()))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString(secret)
	return
}

// ValidateLoginToken 验证登录token
func ValidateLoginToken(signedToken string, secret []byte) (claims *LoginToken, success bool) {
	token, err := jwt.ParseWithClaims(signedToken, &LoginToken{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method %v", token.Header["alg"])
			}
			return secret, nil
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	claims, ok := token.Claims.(*LoginToken)
	if ok && token.Valid {
		success = true
		return
	}
	return
}
