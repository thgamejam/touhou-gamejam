package data

import (
    "context"
    "errors"
    "math/rand"
    "service/app/account/internal/biz"
    "strconv"
)

const (
    lockOpenerListMaxLen = 5
)

var lockOpenerCacheKey = func(hash string) string {
    return "lock_opener_" + hash
}

var lockOpenerIDCacheKey = func(id int) string {
    return "lock_opener_id_to_key_" + strconv.Itoa(id)
}

// GetPublicKey 使用Hash值获取公钥
func (r *accountRepo) GetPublicKey(ctx context.Context, hash string) (*biz.PublicKey, error) {
    lock := &LockOpener{}
    ok, err := r.data.Cache.Get(ctx, lockOpenerCacheKey(hash), lock)
    if err != nil {
        r.log.Error("") // TODO
    }

    if ok {
        return &biz.PublicKey{
            Hash: hash,
            Key:  lock.Public,
        }, nil
    }

    return nil, errors.New("") // TODO
}

// GetRandomlyPublicKey 获取任意的公钥
func (r *accountRepo) GetRandomlyPublicKey(ctx context.Context) (key *biz.PublicKey, err error) {
    // 随机获取id
    id := rand.Intn(lockOpenerListMaxLen)

    // 查找缓存中是否已经存在id对于的密钥对
    var lock *LockOpener
    hash, ok, _ := r.data.Cache.GetString(ctx, lockOpenerIDCacheKey(id))
    if ok {
        return r.GetPublicKey(ctx, hash)
    }

    // 缓存不存在密钥对时进行创建
    lock, hash, err = r.CreateLockOpenerToCache(ctx, id)
    if err != nil {
        return nil, err
    }

    return &biz.PublicKey{
        Hash: hash,
        Key:  lock.Public,
    }, nil
}

// GetPrivateKey 使用Hash值获取密钥
func (r *accountRepo) GetPrivateKey(ctx context.Context, hash string) (*biz.PrivateKey, error) {
    lock := &LockOpener{}
    ok, err := r.data.Cache.Get(ctx, lockOpenerCacheKey(hash), lock)
    if err != nil {
        r.log.Error("") // TODO
    }

    if ok {
        return &biz.PrivateKey{
            Hash: hash,
            Key:  lock.Private,
        }, nil
    }

    return nil, errors.New("") // TODO
}
