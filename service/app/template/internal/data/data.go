package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
	pkgConf "service/pkg/conf"
	"service/pkg/data/cache"
	"service/pkg/data/database"
	"service/pkg/data/object_storage"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// 封装的数据库客户端
	DataBase      *gorm.DB               // 数据库
	Cache         *redis.Client          // 缓存
	ObjectStorage *object_storage.Client // 对象存储服务的封装
}

// NewData .
func NewData(confService *pkgConf.Service, logger log.Logger) (*Data, func(), error) {
	// 数据库
	db, err := database.NewDataBase(confService.Data.Database)
	if err != nil {
		return nil, nil, err
	}
	err = DBAutoMigrate(db)
	if err != nil {
		return nil, nil, err
	}
	// 缓存
	c, err := cache.NewCache(confService.Data.Redis)
	if err != nil {
		return nil, nil, err
	}
	// 对象存储
	os, err := object_storage.NewObjectStorage(confService.Data.ObjectStorage)

	data := &Data{
		DataBase:      db,
		Cache:         c,
		ObjectStorage: os,
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}

// DBAutoMigrate 数据库自动迁移
func DBAutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
	// 添加数据库模型在内
	//&biz.User{},
	)
	if err != nil {
		return err
	}

	return nil
}
