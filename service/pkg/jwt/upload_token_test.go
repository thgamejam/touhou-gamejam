package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	secret       = []byte("61bc667fe31f47f7a312ee177be915dd")
	uploadClaims = UploadClaims{
		FileID:   1,
		MD5:      "0d9ba6459f17d124f211d1762cf17c7e",
		SliceID:  1,
		SliceLen: 1,
	}
)

func TestUploadToken(t *testing.T) {
	token, err := CreateUploadToken(&uploadClaims, secret, time.Hour)
	if err != nil {
		assert.NoError(t, err)
	}
	t.Logf("log token:=%v\n", token)

	claims, success := ValidateUploadToken(token, secret)
	assert.True(t, success)
	assert.NotNil(t, claims)

	assert.Equal(t, uploadClaims.MD5, claims.MD5)

	token, err = CreateUploadToken(&uploadClaims, secret, time.Nanosecond)
	if err != nil {
		assert.NoError(t, err)
	}
	t.Logf("log expires at token:=%v\n", token)

	time.Sleep(time.Second)

	claims, success = ValidateUploadToken(token, secret)
	assert.False(t, success)
	assert.Nil(t, claims)
}
