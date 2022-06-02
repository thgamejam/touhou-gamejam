package jwt

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// LoginToken 登录token
type LoginToken struct {
	UserID     uint32 `json:"user_id"`     // 用户id
	CreateTime int64  `json:"create_time"` // 创建时间
	jwt.RegisteredClaims
}

func ValidateLoginListMatcher() selector.MatchFunc {
	skipRouters := make(map[string]bool)
	skipRouters["/passport.v1.Passport/ChangePassword"] = true
	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return true
		}
		return false
	}
}

func JWTLoginAuth(secret []byte) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				token := tr.RequestHeader().Get("Authorization")
				if token == "" {
					return nil, errors.New("TokenNotFound")
				}
				_, success := ValidateLoginToken(token, secret)
				if !success {
					return nil, errors.New("TokenValidateError")
				}
			}
			return handler(ctx, req)
		}
	}
}

func CreateLoginToken(claims LoginToken, secret []byte, expirationTime time.Duration) (signedToken string, err error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(expirationTime))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString(secret)
	return
}

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
