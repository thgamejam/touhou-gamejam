package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewAccountRepo,
	// TODO 数据客户端构建函数
)

// Data .
type Data struct {
	// TODO 封装的数据客户端
}

// NewData .
func NewData(
	// TODO 需要的数据客户端
	logger log.Logger,
) (*Data, func(), error) {
	data := &Data{}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}
