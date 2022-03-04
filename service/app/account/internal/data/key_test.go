package data

import "testing"

var Key = "-----BEGIN PUBLIC KEY-----\n" +
    "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzDb2y1ybZxeBUcJrb80BCFplXBOQ\n" +
    "sQrDJO03vD78dzIp+N911bq2CJXksVWjqR7zDTkGDF41zoBNOlHSbb04ZA==\n" +
    "-----END PUBLIC KEY-----"

func TestAccount_hashMd5To16(t *testing.T) {
    hash := hashMd5To16(Key)
    t.Logf("log.hash:=%v\n", hash)
}
