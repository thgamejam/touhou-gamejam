package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// UploadClaims 上传文件的token
type UploadClaims struct {
	FileID   uint64 `json:"id"`                  // 文件ID
	MD5      string `json:"md5"`                 // MD5Hash值
	SHA1     string `json:"sha1"`                // SHA1Hash值
	SliceID  uint32 `json:"slice_id,omitempty"`  // 分片ID
	SliceLen uint32 `json:"slice_len,omitempty"` // 分片长度
	jwt.RegisteredClaims
}

// CreateUploadToken 创建上传文件token
func CreateUploadToken(claims *UploadClaims, secret []byte, expirationTime time.Duration) (signedToken string, err error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(expirationTime))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString(secret)
	return
}

// ValidateUploadToken 验证上传文件token
func ValidateUploadToken(signedToken string, secret []byte) (claims *UploadClaims, success bool) {
	token, err := jwt.ParseWithClaims(signedToken, &UploadClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("错误的签名方式 %v", token.Header["alg"])
			}
			return secret, nil
		})
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*UploadClaims)
	if ok && token.Valid {
		success = true
		return
	}

	return
}
