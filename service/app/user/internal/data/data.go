package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	// TODO 数据客户端构建函数
)

// Data .
type Data struct {
	// TODO 封装的数据客户端
	Redis         *redis.Client
	DataBase      *gorm.DB
}

// NewData .
func NewData(
	// TODO 需要的数据客户端
	db *gorm.DB,
	red *redis.Client,
	logger log.Logger,
) (*Data, func(), error) {
	data := &Data{
		DataBase:      db,
		Redis:         red,
	}

	cleanup := func() {
		_ = red.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}
