package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/stretchr/testify/assert"
	v1 "service/api/account/v1"
	"service/pkg/conf"
	pkgConsul "service/pkg/consul"
	"service/pkg/crypto/ecc"
	"testing"
)

const (
	Address          = "localhost:8500"     // Consul 服务器地址
	Scheme           = "http"               // Consul 服务器的 URI 方案 ("http" or "grpc")
	TestEMail        = "test@test.com"      // 测试用邮箱
	TestPasswd       = "test-passwd"        // 测试用密码
	TestChangePasswd = "test-passwd-change" // 测试用修改密码
)

var (
	accountServiceClient = newAccountServiceClient()
	testCtx              = context.Background()
)

func newAccountServiceClient() v1.AccountClient {
	config := conf.Consul{
		Address:    Address,
		Scheme:     Scheme,
		Datacenter: "",
		Path:       "",
	}
	consulUtil := pkgConsul.New(&config)
	rd := consulUtil.NewDiscovery()
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///thjam.account.service"),
		grpc.WithDiscovery(rd),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := v1.NewAccountClient(conn)
	return c
}

func encryptPassword(getKeyReply *v1.GetKeyReply, passwd string) string {
	plaintext := []byte(passwd)
	ecdsaPublicKey, _ := ecc.ParsePublicKey(getKeyReply.Key)
	ciphertext, _ := ecc.Encrypt(ecdsaPublicKey, plaintext)
	return ciphertext
}

func TestGRPCAccount_GetKey(t *testing.T) {
	req := v1.GetKeyReq{}
	key, err := accountServiceClient.GetKey(testCtx, &req)
	assert.NoError(t, err)
	assert.NotEqual(t, key.Key, "", "GetKey返回的 公钥 不应该为 空字符串")
	assert.NotEqual(t, key.Hash, "", "GetKey返回的 公钥哈希值 不应该为 空字符串")
	t.Logf("TestGRPCAccount_GetKey:  GetKey success.")
}

func TestGRPCAccount_CreateEMailAccount(t *testing.T) {
	getKeyReq := v1.GetKeyReq{}
	getKeyReply, _ := accountServiceClient.GetKey(testCtx, &getKeyReq)
	passwd := encryptPassword(getKeyReply, TestPasswd)

	prepareCreateEMailAccountReq := v1.PrepareCreateEMailAccountReq{
		Email:      TestEMail,
		Hash:       getKeyReply.Hash,
		Ciphertext: passwd,
	}
	prepareCreateEMailAccountReply, err := accountServiceClient.PrepareCreateEMailAccount(testCtx, &prepareCreateEMailAccountReq)
	assert.NoError(t, err)
	assert.NotEqual(t, prepareCreateEMailAccountReply.Sid, "", "PrepareCreateEMailAccount返回的 会话号 不应该为 空字符串")
	t.Logf("TestGRPCAccount_CreateEMailAccount:  PrepareCreateEMailAccount success.")

	finishCreateEMailAccountReq := v1.FinishCreateEMailAccountReq{Sid: prepareCreateEMailAccountReply.Sid}
	finishCreateEMailAccountReply, err := accountServiceClient.FinishCreateEMailAccount(testCtx, &finishCreateEMailAccountReq)
	assert.NoError(t, err)
	assert.NotZero(t, finishCreateEMailAccountReply.Id, "FinishCreateEMailAccount返回的 账户id 不应该为 0")
	t.Logf("TestGRPCAccount_CreateEMailAccount:  FinishCreateEMailAccount success.")
}

func TestGRPCAccount_VerifyPassword(t *testing.T) {
	getKeyReq := v1.GetKeyReq{}
	getKeyReply, _ := accountServiceClient.GetKey(testCtx, &getKeyReq)
	passwd := encryptPassword(getKeyReply, TestPasswd)

	req := v1.VerifyPasswordReq{
		Username:   TestEMail,
		Hash:       getKeyReply.Hash,
		Ciphertext: passwd,
	}
	reply, err := accountServiceClient.VerifyPassword(testCtx, &req)
	assert.NoError(t, err)
	assert.NotZero(t, reply.Id, "VerifyPassword返回的 账户id 不应该为 0")
	t.Logf("TestGRPCAccount_VerifyPassword:  VerifyPassword success.")
}

func TestGRPCAccount_GetAccount(t *testing.T) {
	getKeyReq := v1.GetKeyReq{}
	getKeyReply, _ := accountServiceClient.GetKey(testCtx, &getKeyReq)
	passwd := encryptPassword(getKeyReply, TestPasswd)
	req := v1.VerifyPasswordReq{Username: TestEMail, Hash: getKeyReply.Hash, Ciphertext: passwd}
	reply, _ := accountServiceClient.VerifyPassword(testCtx, &req)

	getAccountReq := v1.GetAccountReq{Id: reply.Id}
	getAccountReply, err := accountServiceClient.GetAccount(testCtx, &getAccountReq)
	assert.NoError(t, err)
	assert.NotEqual(t, getAccountReply.Uuid, "", "GetAccount返回的 UUID 不应该为 空字符串")
	assert.NotEqual(t, getAccountReply.Email, "", "GetAccount返回的 邮箱 不应该为 空字符串")
	t.Logf("TestGRPCAccount_GetAccount:  GetAccount success.")
}

func TestGRPCAccount_SavePassword(t *testing.T) {
	getKeyReq := v1.GetKeyReq{}
	getKeyReply, _ := accountServiceClient.GetKey(testCtx, &getKeyReq)
	passwd1 := encryptPassword(getKeyReply, TestPasswd)
	passwd2 := encryptPassword(getKeyReply, TestChangePasswd)
	req := v1.VerifyPasswordReq{Username: TestEMail, Hash: getKeyReply.Hash, Ciphertext: passwd1}
	reply, _ := accountServiceClient.VerifyPassword(testCtx, &req)

	savePasswordReq := v1.SavePasswordReq{
		Id:         reply.Id,
		Ciphertext: passwd2,
		Hash:       getKeyReply.Hash,
	}
	savePasswordReply, err := accountServiceClient.SavePassword(testCtx, &savePasswordReq)
	assert.NoError(t, err)
	_ = savePasswordReply

	req = v1.VerifyPasswordReq{Username: TestEMail, Hash: getKeyReply.Hash, Ciphertext: passwd2}
	reply, err = accountServiceClient.VerifyPassword(testCtx, &req)
	assert.NoError(t, err, "SavePassword修改后的密码不应VerifyPassword验证错误")

	// 复原密码
	savePasswordReq = v1.SavePasswordReq{Id: reply.Id, Ciphertext: passwd1, Hash: getKeyReply.Hash}
	_, err = accountServiceClient.SavePassword(testCtx, &savePasswordReq)
	assert.NoError(t, err, "复原密码不应出现错误")
	t.Logf("TestGRPCAccount_SavePassword:  GetAccount success.")
}
