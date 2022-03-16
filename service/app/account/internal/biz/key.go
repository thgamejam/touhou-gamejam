package biz

import "context"

// PublicKey 公钥
type PublicKey struct {
    Hash string // 密钥摘要
    Key  string // 公钥内容
}

// PrivateKey 密钥
type PrivateKey struct {
    Hash string // 密钥摘要
    Key  string // 密钥内容
}

// GetKey 获取公钥
func (uc *AccountUseCase) GetKey(ctx context.Context, hash string) (*PublicKey, error) {
    return uc.repo.GetPublicKey(ctx, hash)
}

// GetRandomlyKey 获取任意的一个公钥
func (uc *AccountUseCase) GetRandomlyKey(ctx context.Context) (*PublicKey, error) {
    return uc.repo.GetRandomlyPublicKey(ctx)
}
