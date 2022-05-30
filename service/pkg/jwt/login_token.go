package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// LoginToken 登录token
type LoginToken struct {
	UserID     uint32 `json:"user_id"`     // 用户id
	CreateTime int64  `json:"create_time"` // 创建时间
	jwt.RegisteredClaims
}

func CreateLoginToken(claims LoginToken, secret []byte, expirationTime time.Duration) (signedToken string, err error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(expirationTime))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString(secret)
	return
}

func ValidateLoginToken(signedToken string, c jwt.Claims, secret []byte) (claims jwt.Claims, success bool) {
	token, err := jwt.ParseWithClaims(signedToken, c,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method %v", token.Header["alg"])
			}
			return secret, nil
		})
	if err != nil {
		return
	}
	claims, success = token.Claims, true
	return
}
