package ecc

import (
	"service/pkg/uuid"
	"testing"
)

func TestECC_GenerateKey(t *testing.T) {

	privateKeyStr, publicKeyStr, err := GenerateKey()
	if err != nil {
		t.Fatalf("TestECC_GenerateKey:  GenerateKey failed!err:=%v", err)
	}

	t.Logf("TestECC_GenerateKey:  GenerateKey success.")
	t.Logf("\tPrivateKey:\n%v", privateKeyStr)
	t.Logf("\tPublicKey:\n%v", publicKeyStr)
}

func TestECC_ParsePrivateKey(t *testing.T) {

	privateKeyStr, _, _ := GenerateKey()

	_, err := ParsePrivateKey([]byte(privateKeyStr))
	if err != nil {
		t.Fatalf("TestParseECCPrivateKey:  ParsePrivateKey failed!err:=%v", err)
	}

}

func TestECC_EncryptANDDecrypt(t *testing.T) {
	privateKeyStr, publicKeyStr, err := GenerateKey()
	if err != nil {
		t.Fatalf("TestECC_EncryptANDDecrypt:  GenerateKey failed!err:=%v", err)
	}

	ecdsaPublicKey, err := ParsePublicKey([]byte(publicKeyStr))
	if err != nil {
		t.Fatalf("TestECC_EncryptANDDecrypt:  ParsePublicKey failed!err:=%v", err)
	}

	ecdsaPrivateKey, err := ParsePrivateKey([]byte(privateKeyStr))
	if err != nil {
		t.Fatalf("TestECC_EncryptANDDecrypt:  ParsePrivateKey failed!err:=%v", err)
	}

	plaintext := []byte("s@strluck.com")
	t.Logf("TestECC_EncryptANDDecrypt:  log:=plaintext: %v", string(plaintext))

	ciphertext, err := Encrypt(ecdsaPublicKey, plaintext)
	if err != nil {
		t.Fatalf("TestECC:  Encrypt failed!err:=%v", err)
	}
	t.Logf("TestECC_EncryptANDDecrypt:  log:=ciphertext: %v", ciphertext)

	recoveredPlaintext, err := Decrypt(ecdsaPrivateKey, ciphertext)
	if err != nil {
		t.Fatalf("TestECC:  Decrypt failed!err:=%v", err)
	}
	t.Logf("TestECC_EncryptANDDecrypt:  log:=recoveredPlaintext: %v", string(recoveredPlaintext))
}

func BenchmarkECC_Encrypt(b *testing.B) {
	// 停止计时器
	b.StopTimer()

	_, publicKeyStr, err := GenerateKey()
	if err != nil {
		b.Fatalf("BenchmarkECC_Encrypt:  GenerateKey failed!err:=%v", err)
	}

	ecdsaPublicKey, err := ParsePublicKey([]byte(publicKeyStr))
	if err != nil {
		b.Fatalf("BenchmarkECC_Encrypt:  ParsePublicKey failed!err:=%v", err)
	}

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
	if err != nil {
		b.Fatalf("BenchmarkECC_Decrypt:  GenerateKey failed!err:=%v", err)
	}
	ecdsaPublicKey, err := ParsePublicKey([]byte(publicKeyStr))
	if err != nil {
		b.Fatalf("BenchmarkECC_Decrypt:  ParsePublicKey failed!err:=%v", err)
	}
	ecdsaPrivateKey, err := ParsePrivateKey([]byte(privateKeyStr))
	if err != nil {
		b.Fatalf("BenchmarkECC_Decrypt:  ParsePrivateKey failed!err:=%v", err)
	}

	ciphertexts := make([]string, b.N, b.N)

	for i := 0; i < b.N; i++ {
		ciphertext, err := Encrypt(ecdsaPublicKey, []byte(uuid.New().String()))
		if err != nil {
			b.Errorf("BenchmarkECC_Decrypt:  Encrypt failed!err:=%v", err)
		}
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
