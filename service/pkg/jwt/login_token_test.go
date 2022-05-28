package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	token, err := CreateToken(LoginToken{
		UserID:     0,
		CreateTime: 0,
		ExpireTime: 0,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}, []byte("aaa"))
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf(token)
}
