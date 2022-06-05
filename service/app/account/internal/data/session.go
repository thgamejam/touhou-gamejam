package data

import (
	"context"
	"encoding/json"
	"service/pkg/util/strconv"
	"service/pkg/uuid"
)

var accountSessionCacheKey = func(id uint32) string {
	return "account_session_to_id_" + strconv.UItoa(id)
}

func (r *accountRepo) CreateSession(ctx context.Context, id uint32, ip string) (sid string, err error) {
	sid = uuid.New().String()
	cache := AccountSessionCache{IP: ip}
	jsonByte, err := json.Marshal(cache)
	if err != nil {
		return "", err
	}
	jsonString := string(jsonByte)
	err = r.data.Cache.Client.HSet(ctx, accountSessionCacheKey(id), map[string]string{sid: jsonString}).Err()
	if err != nil {
		return "", err
	}
	return
}

func (r *accountRepo) ExistSession(ctx context.Context, id uint32, sid string) (bool, error) {
	bc := r.data.Cache.Client.HExists(ctx, accountSessionCacheKey(id), sid)
	err := bc.Err()
	if err != nil {
		return false, err
	}
	return bc.Val(), nil
}

func (r *accountRepo) CloseSession(ctx context.Context, id uint32, sids ...string) error {
	size := len(sids)
	for i, sid := range sids {
		ok, err := r.ExistSession(ctx, id, sid)
		if err != nil {
			return err
		}
		if !ok {
			sids[i] = ""
			size--
		}
	}
	targetSids := make([]string, size)
	for _, sid := range sids {
		if sid != "" {
			targetSids = append(targetSids, sid)
		}
	}

	return r.data.Cache.Client.HDel(ctx, accountSessionCacheKey(id), targetSids...).Err()
}
