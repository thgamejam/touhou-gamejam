package ecc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

// 秘钥加密签名，公钥解密签名

// 公钥加密密文，私钥解密密文

// GenerateKey 生成ECC公钥/秘钥对
func GenerateKey() (privateKeyStr string, publicKeyStr string, err error) {
	// 生成秘钥
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", err
	}

	// x509编码
	eccPrivateKey, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return "", "", err
	}

	// pem编码
	privateBlock := pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: eccPrivateKey,
	}
	privateKeyStr = string(pem.EncodeToMemory(&privateBlock))

	// 获取公钥
	publicKey := privateKey.PublicKey
	// x509编码
	eccPublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return "", "", err
	}

	// pem编码
	block := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: eccPublicKey,
	}
	publicKeyStr = string(pem.EncodeToMemory(&block))

	return
}

// ParsePrivateKey 解码秘钥
func ParsePrivateKey(data string) (privateKey *ecdsa.PrivateKey, err error) {
	// pem解码
	block, _ := pem.Decode([]byte(data))
	// x509解码
	privateKey, err = x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return
}

// ParsePublicKey 解码共钥
func ParsePublicKey(data string) (publicKey *ecdsa.PublicKey, err error) {
	// pem解密
	block, _ := pem.Decode([]byte(data))
	// x509解密
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey = publicInterface.(*ecdsa.PublicKey)
	return
}

// Encrypt ECC加密
func Encrypt(publicKey *ecdsa.PublicKey, pt []byte) (string, error) {
	// ecdsa key 转换 ecies key
	eciesPublicKey := ecies.ImportECDSAPublic(publicKey)
	// 加密
	ct, err := ecies.Encrypt(rand.Reader, eciesPublicKey, pt, nil, nil)
	if err != nil {
		return "", err
	}
	// base64编码
	ctB64 := base64.StdEncoding.EncodeToString(ct)
	return ctB64, err
}

// Decrypt ECC解密
func Decrypt(privateKey *ecdsa.PrivateKey, ctB64 string) ([]byte, error) {
	// base64解码
	ct, err := base64.StdEncoding.DecodeString(ctB64)
	if err != nil {
		return nil, err
	}
	// ecdsa key 转换 ecies key
	eciesPrivateKey := ecies.ImportECDSA(privateKey)
	// 解密
	m, err := eciesPrivateKey.Decrypt(ct, nil, nil)
	if err != nil {
		return nil, err
	}
	return m, err
}
