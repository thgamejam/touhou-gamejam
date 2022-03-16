package data

import (
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
    "gorm.io/gorm"
    "service/pkg/cache"
    "service/pkg/database"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
    NewData,
    NewAccountRepo,
    cache.NewCache,
    database.NewDataBase,
)

// Data .
type Data struct {
    Cache    *cache.Cache
    DataBase *gorm.DB
}

// NewData .
func NewData(db *gorm.DB, cache *cache.Cache, logger log.Logger) (*Data, func(), error) {
    data := &Data{
        DataBase: db,
        Cache:    cache,
    }

    cleanup := func() {
        _ = cache.Close()
        log.NewHelper(logger).Info("closing the data resources")
    }
    return data, cleanup, nil
}
