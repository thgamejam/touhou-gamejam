package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
	"service/app/user/internal/conf"
	"service/pkg/cache"
	"service/pkg/database"
	"service/pkg/object_storage"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	cache.NewCache,
	database.NewDataBase,
	object_storage.NewObjectStorage,
)

// Data .
type Data struct {
	Cache         *cache.Cache
	DataBase      *gorm.DB
	ObjectStorage *object_storage.ObjectStorage
	Conf          *conf.User
}

// NewData .
func NewData(logger log.Logger, db *gorm.DB, cache *cache.Cache, conf *conf.User, ObjectStorage *object_storage.ObjectStorage) (*Data, func(), error) {
	data := &Data{
		DataBase:      db,
		Cache:         cache,
		ObjectStorage: ObjectStorage,
		Conf:          conf,
	}

	ctx := context.Background()
	ok, err := data.ObjectStorage.ExistBucket(ctx, data.Conf.UserAvatarBucketName)
	if err == nil {
		if !ok {
			err := data.ObjectStorage.CreateBucket(ctx, data.Conf.UserAvatarBucketName)
			if err != nil {
				panic(err)
			}
		}
	}

	cleanup := func() {
		_ = cache.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}
