package ecc

import (
	"github.com/stretchr/testify/assert"
	"service/pkg/uuid"
	"testing"
)

func TestECC_GenerateKey(t *testing.T) {
	privateKeyStr, publicKeyStr, err := GenerateKey()
	assert.NoError(t, err)

	t.Logf("TestECC_GenerateKey:  GenerateKey success.")
	t.Logf("\tPrivateKey:\n%v", privateKeyStr)
	t.Logf("\tPublicKey:\n%v", publicKeyStr)
}

func TestECC_ParsePrivateKey(t *testing.T) {
	privateKeyStr, _, _ := GenerateKey()
	_, err := ParsePrivateKey(privateKeyStr)
	assert.NoError(t, err)
}

func TestECC_EncryptANDDecrypt(t *testing.T) {
	privateKeyStr, publicKeyStr, err := GenerateKey()
	assert.NoError(t, err)

	ecdsaPublicKey, err := ParsePublicKey(publicKeyStr)
	assert.NoError(t, err)

	ecdsaPrivateKey, err := ParsePrivateKey(privateKeyStr)
	assert.NoError(t, err)

	plaintext := []byte("s@strluck.com")
	t.Logf("TestECC_EncryptANDDecrypt:  log:=plaintext: %v", string(plaintext))

	ciphertext, err := Encrypt(ecdsaPublicKey, plaintext)
	t.Logf("TestECC_EncryptANDDecrypt:  log:=plaintext len: %v", len(ciphertext))
	assert.NoError(t, err)
	t.Logf("TestECC_EncryptANDDecrypt:  log:=ciphertext: %v", ciphertext)

	recoveredPlaintext, err := Decrypt(ecdsaPrivateKey, ciphertext)
	assert.NoError(t, err)

	assert.Equal(t, plaintext, recoveredPlaintext)

	t.Logf("TestECC_EncryptANDDecrypt:  log:=recoveredPlaintext: %v", string(recoveredPlaintext))
}

func BenchmarkECC_Encrypt(b *testing.B) {
	// 停止计时器
	b.StopTimer()

	_, publicKeyStr, err := GenerateKey()
	assert.NoError(b, err)

	ecdsaPublicKey, err := ParsePublicKey(publicKeyStr)
	assert.NoError(b, err)

	uuids := make([]string, b.N, b.N)
	for i := 0; i < b.N; i++ {
		uuids[i] = uuid.New().String()
	}

	// 开始计时器
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err := Encrypt(ecdsaPublicKey, []byte(uuids[i]))
		if err != nil {
			b.Errorf("BenchmarkECC_Encrypt:  Encrypt failed!err:=%v", err)
		}
	}
}

func BenchmarkECC_Decrypt(b *testing.B) {
	b.StopTimer()

	privateKeyStr, publicKeyStr, err := GenerateKey()
	assert.NoError(b, err)

	ecdsaPublicKey, err := ParsePublicKey(publicKeyStr)
	assert.NoError(b, err)

	ecdsaPrivateKey, err := ParsePrivateKey(privateKeyStr)
	assert.NoError(b, err)

	ciphertexts := make([]string, b.N, b.N)

	for i := 0; i < b.N; i++ {
		ciphertext, err := Encrypt(ecdsaPublicKey, []byte(uuid.New().String()))
		assert.NoError(b, err)
		ciphertexts[i] = ciphertext
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err := Decrypt(ecdsaPrivateKey, ciphertexts[i])
		if err != nil {
			b.Errorf("BenchmarkECC_Decrypt:  Decrypt failed!err:=%v", err)
		}
	}
}
