package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"service/pkg/cache"
	"service/pkg/object_storage"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewFileEdgeRepo,
	cache.NewCache,
	object_storage.NewObjectStorage,
)

// Data .
type Data struct {
	cache *cache.Cache
	oss   *object_storage.ObjectStorage
}

// NewData .
func NewData(cache *cache.Cache, oss *object_storage.ObjectStorage, logger log.Logger) (*Data, func(), error) {
	data := &Data{
		cache: cache,
		oss:   oss,
	}

	cleanup := func() {
		_ = cache.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}
